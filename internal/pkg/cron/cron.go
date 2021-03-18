package cron

import (
	"vvvstore/internal/app/job"
	"vvvstore/pkg/cron"
)

var CronManger cron.Manager

// 初始化定时任务管理器
func InitCronManger()  {
	CronManger := cron.NewManager()
	CronManger.Register(job.Test{}) // Register Job ...
	CronManger.Start()
}
