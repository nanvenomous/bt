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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	completion string
	shells     []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bt",
	Short: "a package for managing bluetooth devices with bluetoothctl",
	Long:  `a package for managing bluetooth devices with bluetoothctl`,
	Run: func(cmd *cobra.Command, args []string) {
		if completion != "" {
			switch completion {
			case shells[0]:
				cmd.Root().GenBashCompletion(os.Stdout)
			case shells[1]:
				cmd.Root().GenZshCompletion(os.Stdout)
			case shells[2]:
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case shells[3]:
				cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			default:
				fmt.Println("not a recognized shell")
				os.Exit(1)
			}
			os.Exit(0)
		}
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	shells = []string{"bash", "zsh", "fish", "powershell"}
	cobra.OnInitialize(initConfig)

	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	completionFlag := "completion"
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bt.yaml)")
	rootCmd.PersistentFlags().StringVar(&completion, completionFlag, "", "generate shell completion")
	rootCmd.RegisterFlagCompletionFunc(completionFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return shells, cobra.ShellCompDirectiveDefault
	})

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".bt" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bt")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
