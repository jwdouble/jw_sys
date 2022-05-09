package main

import (
	"io"
	"jw.sys/service"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		service.Health(w, &http.Request{})
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/health", h1)
	http.HandleFunc("/endpoint", h2)
	log.Fatal(http.ListenAndServe(":11000", nil))
}
