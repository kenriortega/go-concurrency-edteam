package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var urls = []string{
	"https://github.com/kenriortega/traefik",
	"https://docs.min.io/",
	"https://www.youtube.com/",
}

func main() {
	fetchConcurrent(urls)
}

func fetchConcurrent(urls []string) {
	done := make(chan struct{})

	for _, u := range urls {
		go func(url string) {
			fetch(url)
			select {
			case <-done:
				return
			default:
				fmt.Println("Time exced...", url)
			}
		}(u)
	}

	select {
	case <-time.After(time.Second * 1):
		close(done)
	}
}

func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(url, ": ", resp.StatusCode)
}
