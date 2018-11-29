package main

import (
	"fmt"
)

type result struct {
	int
	string
}

func main() {
	results := make(map[int]string, 3)
	resultChannel := make(chan result)

	urls := []string{
		"http://google.com",
		"http://blog.bab.com",
		"waat://furhurtterwe.geds",
	}

	for i, url := range urls {
		go func(i int, u string) {
			resultChannel <- result{i, u}
		}(i, url)
	}

	for i := 0; i < len(urls); i++ {
		result := <- resultChannel
		results[result.int] = result.string
	}

	fmt.Print(results)
}
