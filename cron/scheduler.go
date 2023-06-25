package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/pkg/errors"
)

// cron格式`*/1 * * * *`，分(0-59) 时(0-23) 日(1-31) 月(1-12) 天(0-6)
const (
	ExpressEveryMin   = "* * * * *"
	ExpressEvery5Min  = "*/5 * * * *"
	ExpressEvery10Min = "*/10 * * * *"
	ExpressEvery30Min = "*/30 * * * *"
	ExpressEveryHour  = "0 * * * *" // 每小时整点执行
	ExpressEveryDay   = "0 1 * * *" // 凌晨1:00执行
	ExpressEveryMonth = "0 1 1 * *" // 每个月
)

// SchedulerCfg Cron调度器配置
type SchedulerCfg struct {
	TimeZone         string // 时区配置，默认为Asia/Shanghai
	Async            bool   // 启动job的方式是阻塞还是非阻塞
	SingletonModeAll bool   // 启动job是否采用单例模式，单例模式下若job如果之前有运行且未完成，则调度器不会重复调度同类型新的job任务(若无特殊要求，推荐开启)
}

// DefaultScheduleCfg 默认Schedule配置
var DefaultScheduleCfg = &SchedulerCfg{
	TimeZone:         "Asia/Shanghai",
	Async:            true,
	SingletonModeAll: true,
}

// Scheduler Cron调度程序
type Scheduler struct {
	cfg   *SchedulerCfg     // 调度器配置
	sched *gocron.Scheduler // gocron 调度器
}

// NewScheduler 初始化一个cron 调度器
func NewScheduler(cfg *SchedulerCfg) (*Scheduler, error) {
	// corn 调度所在时区设定
	if cfg.TimeZone == "" {
		cfg.TimeZone = "Asia/Shanghai"
	}
	timeLoc, err := time.LoadLocation(cfg.TimeZone)
	if err != nil {
		return nil, errors.Wrapf(err, "parse cron timezone Name got err: %s", err)
	}
	s := gocron.NewScheduler(timeLoc)

	// 单例任务模式，即之前有job执行还未结束，则不会开启新的job
	if cfg.SingletonModeAll {
		s.SingletonModeAll()
	}

	return &Scheduler{
		cfg:   cfg,
		sched: s,
	}, nil
}

// AddTasks 新增一个cron task任务
func (s *Scheduler) AddTasks(tasks ...*Task) error {
	for _, task := range tasks {
		cronJob, err := s.sched.Cron(task.Express).Do(task.Execute)
		if err != nil {
			return errors.Wrapf(err, "[error] sched add task[%s] got err", task.Name)
		}
		cronJob.Tag(task.Name)
	}
	return nil
}

// Start 启动Cron定时服务
func (s *Scheduler) Start() {
	// 是否启动非阻塞任务
	if s.cfg.Async {
		s.sched.StartAsync()
	} else {
		s.sched.StartBlocking()
	}
}
