package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// LoggerConfig db日志配置
type LoggerConfig struct {
	LogFile       string          `json:"log_file"`       // 日志存储位置
	LogLevel      logger.LogLevel `json:"log_level"`      // 日志级别， 4:info
	SlowThreshold string          `json:"slow_threshold"` // 慢日志阈值
	Colorful      bool            `json:"colorful"`       // 颜色区分
}

// DBConfig Mysql DB配置
type DBConfig struct {
	// 连接配置
	DSN string `json:"dsn"`
	// 连接池
	ConnMaxLifetime string `json:"conn_max_lifetime"`  // 连接可以被复用多久
	ConnMaxIdleTime string `json:"conn_max_idle_time"` // 连接可以被闲置多久
	MaxIdleConns    int    `json:"max_idle_conns"`     // 连接池最大闲置连接数
	MaxOpenConns    int    `json:"max_open_conns"`     // 连接池最大开启连接数
	// 日志
	LoggerConfig *LoggerConfig `json:"logger"`
}

// String 配置字符串序列化返回，用于信息展示
func (cfg *DBConfig) String() string {
	s, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Sprintf("can not marshal, %s", err)
	}
	return string(s)
}

// NewGormDB 初始MysqlDB 仓储实例
func NewGormDB(cfg *DBConfig) (*gorm.DB, error) {
	// mysql logger
	gormLogger, err := newGormLogger(cfg.LoggerConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "new gorm db fail, newGormLogger() got err")
	}

	// mysql instance
	gormDB, err := gorm.Open(
		mysql.Open(cfg.DSN),
		&gorm.Config{
			Logger: gormLogger,
		},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "new gorm db fail, gorm.Open() got err")
	}

	// mysql connect pool config
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, errors.Wrapf(err, "new gorm db fail, gormDB.DB() got err")
	}

	sqlDB.SetConnMaxLifetime(mustParseDuration(cfg.ConnMaxLifetime))
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	return gormDB, nil
}

// newGormLogger 初始化DB日志记录器
func newGormLogger(logCfg *LoggerConfig) (logger.Interface, error) {
	if logCfg == nil || logCfg.LogFile == "" {
		return nil, nil
	}

	// open logfile
	sqlLog, err := os.OpenFile(logCfg.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, errors.Wrapf(err, "newGormLogger os.OpenFile(%s) fail", logCfg.LogFile)
	}

	newLogger := logger.New(
		log.New(sqlLog, "", log.LstdFlags),
		logger.Config{
			LogLevel:      logCfg.LogLevel,                         // Log level
			SlowThreshold: mustParseDuration(logCfg.SlowThreshold), // 慢SQL阈值
			Colorful:      logCfg.Colorful,                         // 彩色打印
		},
	)
	return newLogger, nil
}

// mustParseDuration 解析配置文件中的字符串时间, 支持单位"ns", "us" (or "µs"), "ms", "s", "m", "h"
func mustParseDuration(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return duration
}
