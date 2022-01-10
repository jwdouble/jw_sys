package cron

import (
	"github.com/robfig/cron/v3"

	"jw.sys/service"
)

func Do() {
	c := cron.New()
	c.AddFunc("0 0 0 * *", service.SysLog)
	c.Run()
}
