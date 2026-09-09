/*
Copyright © 2024 Oliver Schrenk
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

// browserFlag is persistent so `list` and `profiles` share one definition.
var browserFlag string

var rootCmd = &cobra.Command{
	Use:     "cutter",
	Short:   "Extracts cookies from browser",
	Version: version,
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
	rootCmd.PersistentFlags().StringVarP(&browserFlag, "browser", "b", "safari", "Browser (safari|arc)")
}

func initConfig() {
	// read config file here
}
