package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	select {
	case <- ping(a):
		return a
	case <- ping(b):
		return b
	}

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan bool  {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}