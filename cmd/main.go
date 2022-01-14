package main

import (
	"net/http"

	"jw.lib/conf"
	"jw.lib/logx"

	"jw.sys/service"
)

func main() {
	http.HandleFunc("/health", service.Health)
	logx.Error(http.ListenAndServe(conf.GetYaml("app.port"), nil))
}
