package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	c "github.com/oschrenk/cutter/internal"
)

var profileFlag string
var formatFlag string
var domainFlag string

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&profileFlag, "profile", "p", "", "Profile ID (Safari UUID, Chromium profile directory, or 'default')")
	listCmd.Flags().StringVarP(&formatFlag, "format", "f", "json", "Output format (json|netscape)")
	listCmd.Flags().StringVarP(&domainFlag, "domain", "d", "", "Only cookies for this host and its subdomains")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List cookies",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cutter := c.NewInstance()
		cookies := cutter.List(browserFlag, profileFlag)
		cookies = c.FilterDomain(cookies, domainFlag)
		out, err := c.Format(cookies, formatFlag)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out)
	},
}
