package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	c "github.com/oschrenk/cutter/internal"
)

var profileFlag string

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&profileFlag, "profile", "p", "", "Profile ID (Safari UUID, Arc directory, or 'default')")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List cookies",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cutter := c.NewInstance()
		cookies := cutter.List(browserFlag, profileFlag)
		json, err := json.MarshalIndent(cookies, "", "  ")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(json))
		}
	},
}
