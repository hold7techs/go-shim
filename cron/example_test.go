package cron

import (
	"log"
)

// ExampleStart 开始
func ExampleStart() {
	var task1Fn = func(taskID int) {
		log.Printf("Exec task[#%d] fn", taskID)
	}

	// 准备task任务
	task1 := NewTask("task1", func() { task1Fn(1) }, "* * * * *")
	task2 := NewTask("task2", func() { task1Fn(2) }, "*/2 * * * *")

	// 准备scheduler调度器
	// locker, err := NewRedisLocker("redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2")
	// if err != nil {
	// 	log.Fatalf("parse redis url got err: %s", err)
	// }

	crond, _ := NewScheduler(&SchedulerCfg{
		Async:            false, // 不阻塞主协程
		SingletonModeAll: true,  // 调度器不会重复调度同类型新的task任务
	})

	// 添加task任务
	err := crond.AddTasks(task1, task2)
	if err != nil {
		log.Fatalf("crond add task got err: %s", err)
	}

	// task执行
	crond.Start()
}
