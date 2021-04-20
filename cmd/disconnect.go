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
	"github.com/mrgarelli/bt/system"
	"github.com/spf13/cobra"
)

// disconnectCmd represents the disconnect command
var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "<device>: disconnect from a device",
	Long:  `<device>: disconnect from a device`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return system.CheckIfArgumentIsDevice(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceID, err := system.GetDeviceID(system.C.Devices[args[0]])
		if err != nil {
			return err
		}
		err = system.Bluetoothctl([]string{"disconnect", deviceID})
		if err != nil {
			return err
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return system.DeviceNames, cobra.ShellCompDirectiveDefault
	},
}

func init() {
	rootCmd.AddCommand(disconnectCmd)
	system.GetDevicesFromConfig()

	// disconnectCmd.PersistentFlags().String("foo", "", "A help for foo")
	// disconnectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
