package main

import (
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://github.com/kenriortega/traefik",
	"https://docs.min.io/",
	"https://www.youtube.com/",
}

func main() {
	// fetchSequencial(urls)
	fetchConcurrent(urls)
}

func fetchSequencial(urls []string) {
	for _, url := range urls {
		fetch(url)
	}
}
func fetchConcurrent(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, u := range urls {
		go func(url string) {
			fetch(url)
			wg.Done()
		}(u)
	}

	wg.Wait()
}

func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(url, ": ", resp.StatusCode)
}
