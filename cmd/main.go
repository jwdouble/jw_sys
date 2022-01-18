package main

import (
	"net/http"

	"jw.lib/conf"

	"jw.sys/service"
)

func main() {
	go service.LogPush()

	http.HandleFunc("/health", service.Health)

	err := http.ListenAndServe(conf.GetYaml("app.port"), nil)
	if err != nil {
		panic(err)
	}
}
