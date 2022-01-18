package service

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	for {
		select {
		case <-time.Tick(time.Second * 10):
			fmt.Println("logx scan")
			push()
		}
	}
}

func push() {
	cmd := rdx.GetRdxOperator().LRange("logx", 0, -1)

	var list []*logx.Logger

	// 把每条记录都推送到loki上,loki负责持久化 ??? 持久化好像不是很顶
	for _, v := range cmd.Val() {
		l := &logx.Logger{}
		fmt.Println(v)
		err := json.Unmarshal([]byte(v), l)
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
	client := &http.Client{Timeout: 5 * time.Second}
	fmt.Println("logx:", string(buf))
	reader := bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "http://150.158.7.96:23100/loki/api/v1/push", reader)
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
			"level": strconv.Itoa(int(v.Level)),
		}
		var val [][2]string
		val = append(val, [2]string{strconv.Itoa(int(v.CreateAt.UnixNano())), v.Content})

		list[n] = &LokiPushStream{Stream: st, Values: val}
	}

	return list
}
