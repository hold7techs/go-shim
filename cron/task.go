package cron

import (
	"log"
	"time"
)

// Task 待执行的JOB任务
type Task struct {
	Name    string // job描述名称
	Express string // cron格式`*/1 * * * *`，分(0-59) 时(0-23) 日(1-31) 月(1-12) 周天(0-6)
	Exec    func() // job调度后执行的内容

	Locker DistributeLocker // 分布式Redis锁
	TTL    time.Duration    // 分布式锁，过期时间
}

// NewTask 创建一个任务
func NewTask(name string, exec func(), express string) *Task {
	return &Task{
		Name:    name,
		Exec:    exec,
		Express: express,
	}
}

// SetDistributeLocker 设置task任务依赖的分布式锁，以及持锁时间
// 更细粒度的cron任务持有的分布式锁，可以基于task自定义，否则采用默认调度器的统一持锁时间
func (t *Task) SetDistributeLocker(locker DistributeLocker, ttl time.Duration) *Task {
	t.Locker = locker
	t.TTL = ttl
	return t
}

// Execute 任务执行
func (t *Task) Execute() {
	if t.Locker != nil {
		// 抢分布式锁，避免cron多节点同时并发调度执行
		if err := t.Locker.Lock(t.Name, t.TTL); err != nil {
			log.Printf("[cron err] try locker cron task got err: %s", err)
			return
		}

		// 获得锁的执行任务
		log.Printf("[cron ok] try locker cron task ok, now begin exec task [#%s]...", t.Name)
		t.Exec()
		return
	}
}

// SingleExec 单机执行
func (t *Task) SingleExec() func() {
	return func() {
		// 无锁直接执行
		log.Printf("no Locker required task, now begin exec task [#%s]...", t.Name)
		t.Exec()
	}
}
