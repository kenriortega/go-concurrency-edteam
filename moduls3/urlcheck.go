package main

import (
	"log"
	"net/http"
)

var urls = []string{
	"https://github.com/kenriortega/traefik",
	"https://docs.min.io/",
	"https://www.youtube.com/",
}

func main2() {
	fetchConcurrent(urls)
}

func fetchSequencial(urls []string) {
	for _, url := range urls {
		fetch(url)
	}
}
func fetchConcurrent(urls []string) {
	signal := make(chan struct{})

	for _, u := range urls {
		go func(url string) {
			fetch(url)
			signal <- struct{}{}
		}(u)
	}

	<-signal
	<-signal
	<-signal
}

func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(url, ": ", resp.StatusCode)
}
