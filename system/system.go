package system

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"

	"github.com/mrgarelli/kik"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Devices map[string]string
}

var (
	C           conf
	DeviceNames []string
)

func CheckIfArgumentIsDevice(args []string) error {
	if len(args) > 0 {
		for _, dev := range DeviceNames {
			if dev == args[0] {
				return nil
			}
		}
		return errors.New("Invalid Argument: " + args[0])
	}
	return errors.New("connect requires an argument <device>")
}

func GetDevicesFromConfig() {
	usr, err := user.Current()
	kik.FailDebugIf(err, 1)
	confPath := path.Join(usr.HomeDir, ".config", "bt.yaml")
	confFile, err := ioutil.ReadFile(confPath)
	kik.FailDebugIf(err, 1)
	err = yaml.Unmarshal(confFile, &C)
	kik.FailDebugIf(err, 1)

	DeviceNames = make([]string, len(C.Devices))
	i := 0
	for k := range C.Devices {
		DeviceNames[i] = k
		i++
	}
}

func GetDeviceInfo(name string) error {
	fmt.Println(name)
	fullName := C.Devices[name]
	id, err := GetDeviceID(fullName)
	if err != nil {
		return err
	}

	cmd := exec.Command("bash", "-c", "bluetoothctl info "+id+" | egrep 'Name|Connected'")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func GetDeviceID(name string) (id string, err error) {
	bluetoothCmd := exec.Command("bluetoothctl", "devices")
	grepCmd := exec.Command("grep", "-i", name)

	bluetoothStdout, err := bluetoothCmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer bluetoothStdout.Close()

	grepCmd.Stdin = bluetoothStdout
	bluetoothCmd.Start()

	grepOutput, err := grepCmd.Output()
	if err != nil {
		return "", err
	}
	resultID := strings.Split(string(grepOutput), " ")[1]
	return resultID, nil
}

func PrependArgument(arg string, args []string) []string {
	args = append(args, arg)
	args[0], args[len(args)-1] = args[len(args)-1], args[0]
	return args
}

func Bluetoothctl(args []string) error {
	cmd := exec.Command("bluetoothctl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}
