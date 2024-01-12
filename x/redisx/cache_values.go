package redisx

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

// GetString 获取缓存的字符串值，若没有找返回空，若redis失败返回redis操作错误，其他情况正常返回
func (c *Client) GetString(ctx context.Context, format string, injects ...interface{}) (string, error) {
	key := RdKey(format, injects...)
	cmd := c.client.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return cmd.Val(), nil
}
