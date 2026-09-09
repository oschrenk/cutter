package internal

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/browserutils/kooky/browser/chromium"
)

// Chromium derivatives share a profile layout: <user data dir>/<profile>/Cookies,
// with the profile display names in <user data dir>/Local State. They differ in
// the keychain entry holding the cookie encryption key, which is what
// chromium.KeyringConfig selects. kooky's chrome package hardcodes "Chrome Safe
// Storage", so it cannot read the others.
type chromiumBrowser struct {
	userDataDir string
	keyring     *chromium.KeyringConfig
}

// userDataDir is relative to the home directory. Arc nests its profiles under
// "User Data"; Chrome on macOS puts them directly in the application directory.
var chromiumBrowsers = map[string]chromiumBrowser{
	BrowserArc: {
		userDataDir: "/Library/Application Support/Arc/User Data",
		keyring:     chromium.KeyringConfigArc,
	},
	BrowserChrome: {
		userDataDir: "/Library/Application Support/Google/Chrome",
		keyring:     &chromium.KeyringConfig{Browser: BrowserChrome, Account: "Chrome"},
	},
}

// chromiumDefaultProfile is the directory name Chromium gives the first profile.
const chromiumDefaultProfile = "Default"

// localState is the subset of a "Local State" file that names the profiles. The
// keys of info_cache are the profile directory names.
type localState struct {
	Profile struct {
		InfoCache map[string]struct {
			Name string `json:"name"`
		} `json:"info_cache"`
	} `json:"profile"`
}

func (cutter *Cutter) chromiumProfiles(browser string) []Profile {
	dir, _ := os.UserHomeDir()
	root := dir + chromiumBrowsers[browser].userDataDir

	names := map[string]string{}
	stateFile, err := os.ReadFile(filepath.Join(root, "Local State"))
	if err == nil {
		var state localState
		if err := json.Unmarshal(stateFile, &state); err == nil {
			for id, info := range state.Profile.InfoCache {
				names[id] = info.Name
			}
		}
	}

	// The cookie database is the thing cutter reads, so a directory without one
	// is not a profile worth reporting.
	matches, err := filepath.Glob(filepath.Join(root, "*", "Cookies"))
	if err != nil {
		log.Fatal(err)
	}

	profiles := []Profile{}
	for _, match := range matches {
		id := filepath.Base(filepath.Dir(match))
		name := names[id]
		if name == "" {
			name = id
		}
		profiles = append(profiles, Profile{
			Name: name,
			ID:   id,
		})
	}

	return profiles
}

func (cutter *Cutter) chromiumCookies(browser string, profile string) []Cookie {
	dir, _ := os.UserHomeDir()
	config := chromiumBrowsers[browser]

	if profile == "" || profile == "default" {
		profile = chromiumDefaultProfile
	}
	cookiesFile := filepath.Join(dir+config.userDataDir, profile, "Cookies")

	if _, err := os.Stat(cookiesFile); os.IsNotExist(err) {
		log.Fatalf("Profile '%s' not found. Use 'cutter profiles --browser %s' to list available profiles.", profile, browser)
	}

	kookies, err := chromium.ReadCookies(context.Background(), cookiesFile, config.keyring)
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
