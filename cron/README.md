# crond lib

crond包基于`github.com/go-co-op/gocron`的基础上，添加了分布式锁定操作；
基于分布式锁的操作，用以解决在集群模式下，多cron可能在同一时刻发起调度的问题

## 使用

### 单机环境下执行cron调度
```
func ExampleStartTest() {
	var job1Fn = func(jobID int) {
		log.Printf("exec job[#%d] fn", jobID)
	}

	// 准备job任务
	job1 := NewJob("job1", func() { job1Fn(1) }, "* * * * *")
	job2 := NewJob("job2", func() { job1Fn(2) }, "*/2 * * * *")

	// 准备scheduler调度器
	locker, err := NewRedisLocker("redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2")
	if err != nil {
		log.Fatalf("parse redis url got err: %s", err)
	}

	crond := NewScheduler(&Config{
		Async:            false,           // 不阻塞主协程
		SingletonModeAll: true,            // 调度器不会重复调度同类型新的job任务
	}, locker)

	// 添加job任务
	err = crond.AddJobs(job1, job2)
	if err != nil {
		log.Fatalf("crond add job got err: %s", err)
	}

	// job执行
	crond.Start()
}
```