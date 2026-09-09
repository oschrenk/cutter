package internal

import "strings"

// FilterDomain keeps the cookies a browser would send to host and its
// subdomains. Cookie domains are stored three ways for the same site, so
// "youtube.com" has to match "youtube.com", ".youtube.com" and
// "www.youtube.com" alike.
//
// A leading dot on the argument is ignored, and matching is case insensitive,
// because domain names are.
func FilterDomain(cookies []Cookie, host string) []Cookie {
	host = strings.TrimPrefix(strings.ToLower(host), ".")
	if host == "" {
		return cookies
	}
	suffix := "." + host

	filtered := []Cookie{}
	for _, cookie := range cookies {
		domain := strings.ToLower(cookie.Domain)
		if domain == host || strings.HasSuffix(domain, suffix) {
			filtered = append(filtered, cookie)
		}
	}
	return filtered
}
