/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"io/ioutil"
	"os/user"
	"path"
	"strings"

	"github.com/mrgarelli/bt/system"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Devices map[string]string
}

var (
	c           conf
	deviceNames []string
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connec <device>: to to a device" + strings.Join(deviceNames, " "),
	Long:  `connec <device>: to to a device`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			for _, dev := range deviceNames {
				if dev == args[0] {
					return nil
				}
			}
			return errors.New("Invalid Argument: " + args[0])
		}
		return errors.New("connect requires an argument <device>")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// for k := range c.Devices {
		// 	fmt.Println(k)
		// }
		// fmt.Println(deviceNames)
		deviceID, err := system.GetDeviceID(args[0])
		if err != nil {
			return err
		}
		system.Bluetoothctl([]string{"connect", deviceID})
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return deviceNames, cobra.ShellCompDirectiveDefault
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	usr, err := user.Current()
	cobra.CheckErr(err)
	confPath := path.Join(usr.HomeDir, ".config", "bt.yaml")
	confFile, err := ioutil.ReadFile(confPath)
	cobra.CheckErr(err)
	err = yaml.Unmarshal(confFile, &c)
	cobra.CheckErr(err)

	deviceNames = make([]string, len(c.Devices))
	i := 0
	for k := range c.Devices {
		deviceNames[i] = k
		i++
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
