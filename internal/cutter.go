package internal

import (
	"fmt"
	"log"
	"os"
	"strings"

	k "github.com/browserutils/kooky"
	"github.com/browserutils/kooky/browser/chrome"
	"github.com/browserutils/kooky/browser/safari"
)

type Cutter struct {
}

func NewInstance() Cutter {
	return Cutter{}
}

func (cutter *Cutter) List(domain string) []*k.Cookie {
	var cookies []*k.Cookie
	for _, cookie := range cutter.safari() {
		if strings.HasPrefix(cookie.Domain, domain) {
			cookies = append(cookies, cookie)
		}
	}

	return cookies
}

func (cutter *Cutter) safari() []*k.Cookie {
	// "/<USER>"
	dir, _ := os.UserHomeDir()
	// my Arc also has one in "Arc/User Data/Profile 5"
	cookiesFile := dir + "/Library/Containers/com.apple.Safari/Data/Library/Cookies/Cookies.binarycookies"
	cookies, err := safari.ReadCookies(cookiesFile)
	if err != nil {
		log.Fatal(err)
	}
	return cookies
}

func (cutter *Cutter) chrome() {
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
