package main

import (
	"fmt"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

var websites = []string{
	"http://google.com",
	"http://blog.gypsydave5.com",
	"waat://furhurterwe.geds",
}

func myChecker(url string) bool {
	if url != "" {
		return true
	} else {
		return false
	}
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

func main() {
	results := CheckWebsites(myChecker, websites)
	for url, status := range results {
		fmt.Printf("Result for %s: %v", url, status)
		fmt.Println()
	}
}
