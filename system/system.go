package system

import (
	"os"
	"os/exec"
	"strings"
)

// get id
// bluetoothctl devices | grep "Echo Plus-CNX" | awk '{print $2}'

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
