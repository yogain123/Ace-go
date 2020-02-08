package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"https://www.goofg.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/", "https://www.vedantu.com/"}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " is down")
		c <- link
	} else {
		fmt.Println(link + " is working")
		c <- link
	}
}
