package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	c "github.com/oschrenk/cutter/internal"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List cookies",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cutter := c.NewInstance()
		cookies := cutter.List()
		json, err := json.MarshalIndent(cookies, "", "  ")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(json))
		}
	},
}

func init() {
}
