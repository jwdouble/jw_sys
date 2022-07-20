package service

import (
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"jw.lib/jsonx"
	"jw.lib/randx"
	"jw.lib/sqlx"
)

// DataMarker 创建大量数据
func DataMarker() {
	f, _ := os.Create("record.txt")

	stmt, err := sqlx.GetSqlOperator().Prepare(`insert into mock_large (id,data_1,data_2,data_3,data_4,data_5,data_6,data_7,data_8,data_9,data_10) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`)
	if err != nil {
		panic(err)
	}

	type info struct {
		Code int    `json:"code,omitempty"`
		Msg  string `json:"msg,omitempty"`
	}

	var i int32
	ch := make(chan struct{}, 0)
	// print speed trend every min
	go func() {
		c := time.NewTicker(time.Minute)
		for {
			select {
			case <-ch: //
				return
			case <-c.C:
				val := atomic.LoadInt32(&i)
				f.WriteString(strconv.Itoa(int(val)) + "\r\n")
			}
		}
	}()

	data := info{Code: 200, Msg: "msg"}
	current := time.Now()
	for i = 600000; i < 1000000; atomic.AddInt32(&i, 1) {
		n := time.Now().UnixNano()
		_, err = stmt.Exec(i, "const", "prefix"+randx.NewString(8, n, 1), 0, randx.NewInt(32, n, 2), randx.NewString(10, n, 3), randx.NewString(20, n, 4), jsonx.MustMarshal(&data), randx.NewString(20, n, 5), randx.NewString(20, n, 6), randx.NewString(20, n, 7))
		if err != nil {
			panic(err)
		}
	}

	dur := time.Since(current)
	fmt.Println("insert finish", dur)
	f.Write([]byte(dur.String()))
	ch <- struct{}{}
}
