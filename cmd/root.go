/*
Copyright © 2024 Oliver Schrenk
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cutter",
	Short: "Extracts cookies from browser",
}

func Execute() {
	// hide (but not disable) "completion" feature
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// read config file here
}
