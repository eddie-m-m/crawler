package main

import "fmt"

func main() {
  normalizedURL, err := NormalizeURL("https://blog.boot.dev/path/")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(normalizedURL)
  }
}
