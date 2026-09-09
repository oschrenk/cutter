package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	c "github.com/oschrenk/cutter/internal"
)

func init() {
	rootCmd.AddCommand(profilesCmd)
}

var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List browser profiles",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cutter := c.NewInstance()
		profiles := cutter.Profiles(browserFlag)
		json, err := json.MarshalIndent(profiles, "", "  ")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(json))
		}
	},
}
