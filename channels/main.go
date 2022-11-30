package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://notexist.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkLink(c, url)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(c, link)
		}(l)
	}
}

func checkLink(c chan string, url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("link %q might be down\n", url)
		c <- url
		return
	}
	fmt.Printf("link %q is up\n", url)
	c <- url
}
