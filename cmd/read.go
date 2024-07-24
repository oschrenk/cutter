package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/browserutils/kooky/browser/chrome"
	"github.com/browserutils/kooky/browser/safari"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read cookies",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		readSafari()
	},
}

func readChrome() {
	fmt.Println("Chrome")
	// "/<USER>/Library/Application Support"
	dir, _ := os.UserConfigDir()
	// my Arc also has one in "Arc/User Data/Profile 5"
	cookiesFile := dir + "/Arc/User Data/Profile 5/Cookies"
	cookies, err := chrome.ReadCookies(cookiesFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}

func readSafari() {
	fmt.Println("Safari")
	// "/<USER>"
	dir, _ := os.UserHomeDir()
	// my Arc also has one in "Arc/User Data/Profile 5"
	cookiesFile := dir + "/Library/Containers/com.apple.Safari/Data/Library/Cookies/Cookies.binarycookies"
	cookies, err := safari.ReadCookies(cookiesFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}
