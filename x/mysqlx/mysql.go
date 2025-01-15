package mysqlx

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
	LogFile       string          `json:"log_file" yaml:"log_file"`             // 日志存储位置
	LogLevel      logger.LogLevel `json:"log_level" yaml:"log_level"`           // 日志级别， 4:info
	SlowThreshold string          `json:"slow_threshold" yaml:"slow_threshold"` // 慢日志阈值
	Colorful      bool            `json:"colorful" yaml:"colorful"`             // 颜色区分
}

// DBConfig Mysql DB配置
type DBConfig struct {
	// 连接配置
	DSN string `json:"dsn" yaml:"dsn"`

	// 连接池
	ConnMaxLifetime string `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`   // 连接可以被复用多久
	ConnMaxIdleTime string `json:"conn_max_idle_time" yaml:"conn_max_idle_time"` // 连接可以被闲置多久
	MaxIdleConns    int    `json:"max_idle_conns" yaml:"max_idle_conns"`         // 连接池最大闲置连接数
	MaxOpenConns    int    `json:"max_open_conns" yaml:"max_open_conns"`         // 连接池最大开启连接数

	// 日志
	LoggerConfig *LoggerConfig `json:"logger" yaml:"logger_config"`
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
func NewGormDB(dbCfg *DBConfig) (*gorm.DB, error) {
	if dbCfg == nil {
		return nil, errors.New("dbCfg is nil")
	}

	// mysql logger
	gormLogger, err := newGormLogger(dbCfg.LoggerConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "newGormLogger() got err")
	}

	// mysql instance
	gormDB, err := gorm.Open(
		mysql.Open(dbCfg.DSN),
		&gorm.Config{
			Logger: gormLogger,
		},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "gorm.Open() got err")
	}

	// mysql connect pool config
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, errors.Wrapf(err, "gormDB.DB() got err")
	}

	sqlDB.SetConnMaxLifetime(mustParseDuration(dbCfg.ConnMaxLifetime))
	sqlDB.SetConnMaxIdleTime(mustParseDuration(dbCfg.ConnMaxIdleTime))
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)

	return gormDB, nil
}

// newGormLogger 初始化DB日志记录器
func newGormLogger(logCfg *LoggerConfig) (logger.Interface, error) {
	if logCfg == nil || logCfg.LogFile == "" || logCfg.LogFile == "console" {
		return logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				LogLevel:      logger.Info,            // Log level
				SlowThreshold: 200 * time.Millisecond, // 慢SQL阈值
				Colorful:      true,                   // 彩色打印
			},
		), nil
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
		log.Fatalf("ParseDuration(%s) got err: %s", s, err)
	}
	return duration
}
