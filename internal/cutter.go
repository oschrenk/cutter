package internal

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/browserutils/kooky/browser/safari"
	"github.com/go-sqlite/sqlite3"
)

type Cutter struct {
}

func NewInstance() Cutter {
	return Cutter{}
}

// Browser identifiers accepted by List and Profiles.
const (
	BrowserSafari = "safari"
	BrowserArc    = "arc"
	BrowserChrome = "chrome"
)

func (cutter *Cutter) List(browser string, profile string) []Cookie {
	switch browser {
	case "", BrowserSafari:
		return cutter.safari(profile)
	case BrowserArc, BrowserChrome:
		return cutter.chromiumCookies(browser, profile)
	default:
		log.Fatalf("Unknown browser '%s'. Use '%s', '%s' or '%s'.", browser, BrowserSafari, BrowserArc, BrowserChrome)
		return nil
	}
}

func (cutter *Cutter) Profiles(browser string) []Profile {
	switch browser {
	case "", BrowserSafari:
		return cutter.safariProfiles()
	case BrowserArc, BrowserChrome:
		return cutter.chromiumProfiles(browser)
	default:
		log.Fatalf("Unknown browser '%s'. Use '%s', '%s' or '%s'.", browser, BrowserSafari, BrowserArc, BrowserChrome)
		return nil
	}
}

// Column indices for bookmarks table (id is included at position 0)
const (
	// colID           = 0
	// colSpecialID    = 1
	// colParent       = 2
	colType  = 3
	colTitle = 4
	// colURL          = 5
	// colNumChildren  = 6
	// colEditable     = 7
	// colDeletable    = 8
	// colHidden       = 9
	// colHiddenAncCnt = 10
	// colOrderIndex   = 11
	colExternalUUID = 12
	// ... columns 13-31 ...
	colSubtype = 32
)

// Bookmark type and subtype values for profiles
const (
	bookmarkTypeFolder     = 1
	bookmarkSubtypeProfile = 2
)

func (cutter *Cutter) safariProfiles() []Profile {
	dir, _ := os.UserHomeDir()
	dbPath := dir + "/Library/Containers/com.apple.Safari/Data/Library/Safari/SafariTabs.db"

	db, err := sqlite3.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	profiles := []Profile{}
	err = db.VisitTableRecords("bookmarks", func(rowId *int64, rec sqlite3.Record) error {
		if len(rec.Values) <= colSubtype {
			return nil
		}

		typeVal, ok := rec.Values[colType].(int)
		if !ok || typeVal != bookmarkTypeFolder {
			return nil
		}

		// subtype can be int or int8 depending on storage
		var subtypeVal int
		switch v := rec.Values[colSubtype].(type) {
		case int:
			subtypeVal = v
		case int8:
			subtypeVal = int(v)
		default:
			return nil
		}
		if subtypeVal != bookmarkSubtypeProfile {
			return nil
		}

		title, _ := rec.Values[colTitle].(string)
		externalUUID, _ := rec.Values[colExternalUUID].(string)

		name := title
		if externalUUID == "DefaultProfile" {
			name = "Default"
		}

		profiles = append(profiles, Profile{
			Name: name,
			ID:   externalUUID,
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return profiles
}

func (cutter *Cutter) safari(profile string) []Cookie {
	dir, _ := os.UserHomeDir()

	var cookiesFile string
	if profile == "" || profile == "default" || profile == "DefaultProfile" {
		cookiesFile = dir + "/Library/Containers/com.apple.Safari/Data/Library/Cookies/Cookies.binarycookies"
	} else {
		// Profile UUIDs are stored lowercase in filesystem
		uuid := strings.ToLower(profile)
		cookiesFile = dir + "/Library/Containers/com.apple.Safari/Data/Library/WebKit/WebsiteDataStore/" + uuid + "/Cookies/Cookies.binarycookies"
	}

	if _, err := os.Stat(cookiesFile); os.IsNotExist(err) {
		log.Fatalf("Profile '%s' not found. Use 'cutter profiles' to list available profiles.", profile)
	}

	kookies, err := safari.ReadCookies(context.Background(), cookiesFile)
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
