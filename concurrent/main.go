package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func returnContentType(url string) string {
	rs, err := http.Get(url)
	if err != nil {
		log.Printf("error: %s\n", err)
		return ""
	}
	defer rs.Body.Close()

	ctype := rs.Header.Get("content-type")
	return ctype
}

func printAll(urls []string) {
	for _, url := range urls {
		log.Printf("The content-type is: %s\n", returnContentType(url))
	}
}

func printAllWithGoRoutine(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			log.Printf("The content-type is: %s\n", returnContentType(url))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func printAllWithChannels(urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			ch <- returnContentType(url)
		}(url)
	}
	close(ch)

	for rs := range ch {
		log.Printf("The content-type is: %s\n", rs)
	}
}

func main() {
	fmt.Printf("The Go Concurrent!")

	urls := []string{"https://golang.org", "https://api.github.com", "https://httpbin.org/xml"}

	printAll(urls)              // real 0m1.925s
	printAllWithGoRoutine(urls) // real 0m1.530s
	printAllWithChannels(urls)  // real 0m1.014s
}
