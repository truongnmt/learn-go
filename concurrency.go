package main

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]string  {
	results := make(map[string]string)

	urls2 := []string{
		"http://google.com",
		"http://blog.bab.com",
		"waat://furhurtterwe.geds",
	}

	for _, url := range urls2 {
		go func() {
			results[url] = url
		}()
	}

	time.Sleep(2 * time.Second)

	return results
}
