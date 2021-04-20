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

// beDiscoverableCmd represents the beDiscoverable command
var beDiscoverableCmd = &cobra.Command{
	Use:   "beDiscoverable",
	Short: "make the current device discoverable",
	Long:  `make the current device discoverable`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := system.Bluetoothctl([]string{"pairable", "on"})
		if err != nil {
			return err
		}
		err = system.Bluetoothctl([]string{"discoverable", "on"})
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(beDiscoverableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// beDiscoverableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// beDiscoverableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
