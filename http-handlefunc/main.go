package main

import (
	"io"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world,this is version 1.")
}

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
