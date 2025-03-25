package main

import (
	"errors"
	"net/url"
	"path"
)


func NormalizeURL(URL string) (string, error) {
  rawURL, err := url.Parse(URL)
  if err != nil {
    return "", errors.New("url cannot be parsed") 
  }
  newURL := rawURL.Hostname() + path.Clean(rawURL.Path)

  return newURL, nil
}
