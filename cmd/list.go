package cmd

import (
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
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cutter := c.NewInstance()
		for _, cookie := range cutter.List("movementsyoga.com") {
			fmt.Println(cookie.Domain, cookie.Name, cookie.Value)
		}
	},
}
