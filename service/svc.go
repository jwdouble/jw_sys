package service

import (
	"encoding/json"
	"net/http"
	"time"

	"jw.lib/conf"
	"jw.lib/logx"
	"jw.lib/rdx"
)

type serverInfo struct {
	Name    string `json:"name,omitempty"`
	Port    string `json:"port,omitempty"`
	Version string `json:"version,omitempty"`
	RunTime int    `json:"runTime,omitempty"`
}

var ServerInfo = serverInfo{
	Name:    conf.GetYaml("app.name"),
	Port:    conf.GetYaml("app.port"),
	Version: conf.GetYaml("app.version"),
}

func init() {
	rdx.Register(conf.AppRedisConn.Value(rdx.DefaultRedisAddr))
	startAt()
}

func startAt() {
	rdx.GetRdxOperator().Set(ServerInfo.Name+"StartAt", time.Now().Unix(), time.Hour*24*14)
}

func Health(w http.ResponseWriter, r *http.Request) {
	buf, _ := json.Marshal(ServerInfo)
	_, err := w.Write(buf)
	if err != nil {
		logx.Error(err)
	}

	return
}
