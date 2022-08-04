package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://google.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}

// contentType will return the value of Content-Type header returned by making
// an HTTP GET request to url
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // https://stackoverflow.com/q/18598780

	// Panic if there's no Content-Type header
	// return resp.Header["Content-Type"][0], nil

	// https://pkg.go.dev/net/http#Header.Get
	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("No Content-Type header!")
	}
	return ctype, nil
}
