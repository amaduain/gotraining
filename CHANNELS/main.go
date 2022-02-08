package main

import (
	"fmt"
	"net/http"
	"time"
)

type counters map[string]int

func main() {
	links := []string{
		"http://www.google.com",
		"http://amazon.com",
		"http://www.facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}
	linkCounters := counters{
		"http://www.google.com":    0,
		"http://amazon.com":        0,
		"http://www.facebook.com":  0,
		"http://stackoverflow.com": 0,
		"http://golang.org":        0,
	}
	c := make(chan string)
	for _, link := range links {
		linkCounters[link] = linkCounters[link] + 1
		fmt.Println(link, "Check #:", linkCounters[link])
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			linkCounters[link] = linkCounters[link] + 1
			fmt.Println(link, "Check #:", linkCounters[link])
			checkLink(link, c)
		}(l)

	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not available!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}
