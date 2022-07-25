package service

import (
	"log"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"jw.lib/jsonx"
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
	rdx.Register(rdx.RedisConfigMap)
	for {
		select {
		case <-time.Tick(time.Second * 10):
			push()
		}
	}
}

func push() {
	pushData := &LokiPushReq{}
	// 把每条记录都推送到loki上,loki负责持久化 ??? 持久化好像不是很顶
	tn := time.Now().UnixNano()
	for {
		val := rdx.GetRdxOperator().LPop("logx")
		if val.Val() == "" {
			break
		}

		i := info{}
		m := map[string]string{
			"app": "uncategorized",
		}
		err := jsonx.Unmarshal([]byte(val.Val()), &i)
		if err == nil {
			m["app"] = i.App
		}

		lps := &LokiPushStream{
			Stream: m,
			Values: [][2]string{
				{strconv.Itoa(int(tn)), val.Val()},
			},
		}

		pushData.Streams = append(pushData.Streams, lps)
		time.Sleep(10 * time.Millisecond)
	}

	if len(pushData.Streams) == 0 {
		return
	}

	_, err := resty.New().SetDebug(true).
		R().SetHeader("Content-Type", "application/json").
		SetBody(jsonx.MustMarshal(pushData)).Post("http://150.158.7.96:3100/loki/api/v1/push")
	if err != nil {
		log.Fatalln("http post err", err.Error())
	}
}

type info struct {
	App string `json:"app,omitempty"`
}

//{
//    "streams": [
//        {
//            "stream": {
//                // 标签，标签内容  -- 用于查询
//                "app": "appName"
//            },
//            "values": [
//                [   // 时间
//                    "1658717309285135071",
//                    // 内容
//                    "fizzbuzz"
//                ]
//            ]
//        }
//    ]
//}
