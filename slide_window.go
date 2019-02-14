package main

import (
	"fmt"
	"net/http"
)

type counter struct {
	c uint64
}

func (cnt *counter) showTotalRequest(w http.ResponseWriter, r *http.Request) {
	cnt.c += 1
	message := fmt.Sprintf("Total request: %d", cnt.c)
	w.Write([]byte(message))
	fmt.Println(message)
}

func showTotalRequest(counter uint64) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		counter += 1
		message := fmt.Sprintf("Total request: %d", counter)
		writer.Write([]byte(message))
		fmt.Println(message)
	}
}

func main() {
	//var c chan  int
	c := make(chan int)

	go func(){
		cnt := counter{c: 0}
		http.HandleFunc("/o", cnt.showTotalRequest)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	go func() {
		http.HandleFunc("/f", showTotalRequest(0))
		if err := http.ListenAndServe(":8888", nil); err != nil {
			panic(err)
		}
	}()

	<- c
}