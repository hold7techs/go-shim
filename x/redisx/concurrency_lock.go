package redisx

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

// ICurrencyLocker 分布式并发锁接口，支持通过业务操作唯一Key进行短时间锁定
type ICurrencyLocker interface {
	DistrbLock(ctx context.Context, lockKey string, ttl time.Duration) error
	DistrbUnLock(ctx context.Context, busKey string) error
}

// DistrbLock 通过Redis增加分布式记录锁设计，防止业务操作短时并发
func (c *Client) DistrbLock(ctx context.Context, lockKey string, ttl time.Duration) error {
	if !c.client.SetNX(ctx, lockKey, true, ttl).Val() {
		return errors.New("concurrency lock crash")
	}
	return nil
}

// DistrbUnLock 分布式解锁操作
func (c *Client) DistrbUnLock(ctx context.Context, busKey string) error {
	return c.client.Del(ctx, busKey).Err()
}
