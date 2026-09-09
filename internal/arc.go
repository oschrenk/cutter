package internal

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/browserutils/kooky/browser/chromium"
)

// Arc is a Chromium derivative. It keeps a standard Chromium profile tree but
// encrypts cookie values with a key stored under its own keychain entry, "Arc
// Safe Storage". kooky's chrome package hardcodes "Chrome Safe Storage", so the
// chromium package plus KeyringConfigArc is what reads Arc correctly.
const arcUserDataDir = "/Library/Application Support/Arc/User Data"

// arcDefaultProfile is the directory name Chromium gives the first profile.
const arcDefaultProfile = "Default"

// localState is the subset of Arc's "Local State" file that names the profiles.
// The keys of info_cache are the profile directory names.
type localState struct {
	Profile struct {
		InfoCache map[string]struct {
			Name string `json:"name"`
		} `json:"info_cache"`
	} `json:"profile"`
}

func (cutter *Cutter) arcProfiles() []Profile {
	dir, _ := os.UserHomeDir()
	root := dir + arcUserDataDir

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

func (cutter *Cutter) arc(profile string) []Cookie {
	dir, _ := os.UserHomeDir()

	if profile == "" || profile == "default" {
		profile = arcDefaultProfile
	}
	cookiesFile := filepath.Join(dir+arcUserDataDir, profile, "Cookies")

	if _, err := os.Stat(cookiesFile); os.IsNotExist(err) {
		log.Fatalf("Profile '%s' not found. Use 'cutter profiles --browser arc' to list available profiles.", profile)
	}

	kookies, err := chromium.ReadCookies(context.Background(), cookiesFile, chromium.KeyringConfigArc)
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
