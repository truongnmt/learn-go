package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	mockWebsiteChecker := func (url string) bool {
		fmt.Println("url: -----        ", url)
		if url == "waat://furhurtterwe.geds" {
			return false
		}
		return true
	}

	websites := []string{
		"http://google.com",
		"http://blog.bab.com",
		"waat://furhurtterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":        true,
		"http://blog.bab.com":      true,
		"waat://furhurtterwe.geds": false,
	}

	t.Run("test label", func(t *testing.T) {
		got := CheckWebsites(mockWebsiteChecker, websites)
		fmt.Println("got -------", got)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
