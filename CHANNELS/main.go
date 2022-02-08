package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://amazon.com",
		"http://www.facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not available!")
		return
	}
	fmt.Println(link, "is up!")
}
