package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	ch <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
}
