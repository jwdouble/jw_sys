package main

import (
	"io"
	"log"
	"net/http"

	"jw.sys/service"
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

	//go service.LogPush()
	go service.DataMarker()

	log.Fatal(http.ListenAndServe(":10001", nil))
}
