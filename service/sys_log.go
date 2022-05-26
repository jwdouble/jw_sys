package service

import (
	"bytes"
	"encoding/json"
	"jw.lib/conf"
	"net/http"
	"strconv"
	"time"

	"jw.lib/logx"
	"jw.lib/rdx"
)

type LokiPushReq struct {
	Streams []*LokiPushStream `json:"streams"`
}

type LokiPushStream struct {
	Stream map[string]string `json:"stream"`
	Values [][2]string       `json:"values"` // "<unix epoch in nanoseconds>", "<log line>"
}

// LogPush 日志信息存在redis中，并推送一份到loki。redis做持久化保存
func LogPush() {
	rdx.Register(conf.APP_REDIS_ADDR.Value(rdx.DefaultRedisAddr), rdx.RedisPwd)
	for {
		select {
		case <-time.Tick(time.Second * 10):
			push()
		}
	}
}

func push() {
	var list []*logx.Logger

	// 把每条记录都推送到loki上,loki负责持久化 ??? 持久化好像不是很顶
	for {

		sc := rdx.GetRdxOperator().LPop("logx")
		if sc.String() == "lpop logx: redis: nil" {
			break
		}
		l := &logx.Logger{}
		err := json.Unmarshal([]byte(sc.Val()), l)
		if err != nil {
			panic(err)
		}
		list = append(list, l)
	}

	streams := LokiPushReq{Streams: parseIn(list)}
	buf, err := json.Marshal(streams)
	if err != nil {
		panic(err)
	}
	client := &http.Client{Timeout: 3 * time.Second}
	reader := bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "http://150.158.7.96:3100/loki/api/v1/push", reader)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func parseIn(l []*logx.Logger) []*LokiPushStream {
	list := make([]*LokiPushStream, len(l))

	for n, v := range l {
		st := map[string]string{
			"level":    v.Level.String(),
			"file":     v.Position,
			"funcName": v.FuncName,
		}
		var val [][2]string
		val = append(val, [2]string{strconv.Itoa(int(v.CreateAt.UnixNano())), logFormat(v)})

		list[n] = &LokiPushStream{Stream: st, Values: val}
	}

	return list
}

func logFormat(l *logx.Logger) string {
	return "FuncName:" + l.FuncName + "	" + "Content:" + l.Content + "\r\n" + "Position:" + l.Position
}
