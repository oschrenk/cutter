package internal

import (
	"log"
	"os"

	"github.com/browserutils/kooky/browser/safari"
)

type Cutter struct {
}

func NewInstance() Cutter {
	return Cutter{}
}

func (cutter *Cutter) List() []Cookie {
	return cutter.safari()
}

func (cutter *Cutter) safari() []Cookie {
	dir, _ := os.UserHomeDir()
	cookiesFile := dir + "/Library/Containers/com.apple.Safari/Data/Library/Cookies/Cookies.binarycookies"
	kookies, err := safari.ReadCookies(cookiesFile)
	if err != nil {
		log.Fatal(err)
	}
	cookies := []Cookie{}
	for _, kookie := range kookies {
		cookies = append(cookies, Cookie{
			kookie.Name,
			kookie.Value,
			kookie.Path,
			kookie.Domain,
			kookie.Expires,
			kookie.MaxAge,
			kookie.Secure,
			kookie.HttpOnly,
			kookie.Creation,
		})
	}

	return cookies
}
