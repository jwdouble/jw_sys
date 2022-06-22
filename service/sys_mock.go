package service

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"jw.lib/conf"
	"jw.lib/jsonx"
	"jw.lib/randx"
	"jw.lib/sqlx"
)

// DataMarker 创建大量数据
func DataMarker() {
	sqlx.Register(sqlx.DefaultSqlDriver, conf.APP_PG_ADDR.Value(sqlx.DefaultSqlAddr))
	f, _ := os.Create("record.txt")

	stmt, err := sqlx.GetSqlOperator().Prepare(`insert into mock_large (id,data_1,data_2,data_3,data_4,data_5,data_6,data_7,data_8,data_9,data_10) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`)
	if err != nil {
		panic(err)
	}

	type info struct {
		Code int    `json:"code,omitempty"`
		Msg  string `json:"msg,omitempty"`
	}

	i := 0
	ch := make(chan struct{}, 0)
	// print speed trend every min
	go func() {
		mu := sync.Mutex{}
		c := time.NewTicker(time.Minute)
		for {
			select {
			case <-ch: //
				return
			case <-c.C:
				mu.Lock()
				f.WriteString(strconv.Itoa(i) + "\r\n")
				mu.Unlock()
			}
		}
	}()

	data := info{Code: 200, Msg: "msg"}
	current := time.Now()
	for i = 0; i < 200000; i++ {
		_, err = stmt.Exec(i, "const", "prefix"+randx.NewString(8), 0, randx.NewInt(32), randx.NewString(10), randx.NewString(20), jsonx.MustMarshal(&data), randx.NewString(20), randx.NewString(20), randx.NewString(20))
		if err != nil {
			panic(err)
		}
	}

	dur := time.Since(current)
	fmt.Println("insert finish", dur)
	f.Write([]byte(dur.String()))
	ch <- struct{}{}
}
