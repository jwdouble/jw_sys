package service

import (
	"encoding/json"
	"jw.lib/logx"
	"net/http"
	"time"
)

type serverInfo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	RunTime string `json:"runTime,omitempty"`
}

var si = serverInfo{
	Name: "jw-sys",
}

func init() {
	//rdx.Register(conf.AppRedisConn.Value(rdx.DefaultRedisAddr))
	//startAt()
	si.RunTime = time.Now().Format("2006-01-02 15:04:05")
}

//func startAt() {
//	rdx.GetRdxOperator().Set(ServerInfo.Name+"StartAt", time.Now().Unix(), time.Hour*24*14)
//}

func Health(w http.ResponseWriter, r *http.Request) {
	buf, _ := json.Marshal(si)
	_, err := w.Write(buf)
	if err != nil {
		logx.Error(err)
	}

	return
}
