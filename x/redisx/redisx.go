package redisx

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=./cache.go -package rdshand -destination cache_mock.go

// IRedisHand redis interface
type IRedisHand interface {
	// GetClient 获取实例
	GetClient() *redis.Client

	// PipeWrite 管道写
	PipeWrite(ctx context.Context, objs []interface{}, kfn KeyFunc, ttl time.Duration) error

	// PipeRead 管道读取
	PipeRead(ctx context.Context, keys []string) (map[string]string, error)

	// GetString 按格式获取
	GetString(ctx context.Context, keyFormat string, injects ...interface{}) (string, error)

	// DelValues 通过key批量删除值
	DelValues(ctx context.Context, keys ...string) error
}

// Client RedisCache
type Client struct {
	client *redis.Client
}

// New create redis obj
func New(client *redis.Client) *Client {
	if client == nil {
		log.Fatal("New nil redis.Client")
	}
	return &Client{client: client}
}

// GetClient get client from Client obj
func (c *Client) GetClient() *redis.Client {
	return c.client
}
