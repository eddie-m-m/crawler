package main

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("cannot parse URL: %w", err)
	}

	newURL := parsedURL.Hostname() + parsedURL.Path
  newURL = strings.ToLower(newURL)
  newURL = strings.TrimSuffix(newURL, "/")

	return newURL, nil
}
