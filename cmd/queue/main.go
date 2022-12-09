package main

import (
	"gin/cmd/queue/job"
	"gin/core"
	"gin/core/redis"
	"github.com/reugn/go-quartz/quartz"
	"time"
)

func main() {
	core.InitDb()
	redis.InitRedis()
	QueueJobs()
}

func QueueJobs() {

	sched := quartz.NewStdScheduler()

	sched.Start()

	exitChan := make(chan bool, 1)

	cronTrigger := quartz.NewRunOnceTrigger(time.Second * 1)

	cronJob := job.Queue{"简单队列消费", exitChan}

	sched.ScheduleJob(&cronJob, cronTrigger)

	for i := 0; i < 1; i++ {
		<-exitChan
	}

	sched.Stop()
}
