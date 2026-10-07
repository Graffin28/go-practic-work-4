package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.reddit.com",
		"https://www.medium.com",
		"https://www.stackoverflow.com",
		"https://www.quora.com",
		"https://www.wikipedia.org",
	}

	jobs := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go worker(jobs, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)
	wg.Wait()
}

func worker(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		response.Body.Close()
		fmt.Println(response.Status, url)
	}
}
