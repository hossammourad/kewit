package helpers

import (
	"net/url"
)

func IsValidURL(u string) bool {
	parsedUrl, err := url.ParseRequestURI(u)
	if err != nil || parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return false
	}
	return true
}
