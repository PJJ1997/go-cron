package cron

import (
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	addFunc()
	addJob()
}

func addFunc() {
	c := cron.New()
	spec := "*/5 * * * * ?" // 每5秒执行一次
	// spec := "0 */5 * * * ?"  每5分钟执行一次
	// spec := "0 0 5 * * ?"  每天5点执行一次
	// spec := "0 0 5 5 * ?"  每月5号5点执行一次
	// spec := "0 5,6,7 * * * ?"  每5分，6分，7分执行一次
	c.AddFunc(spec, func() {
		fmt.Println("task done")
	})
	c.Start()
	select {}
}

type testJob struct{}

func (t testJob) Run() {
	fmt.Println("task done")
}

func addJob() {
	c := cron.New()
	c.AddJob("@yearly", testJob{})
	c.AddJob("@monthly", testJob{})
	c.AddJob("@weekly", testJob{})
	c.AddJob("@daily", testJob{})
	c.AddJob("@hourly", testJob{})
}
