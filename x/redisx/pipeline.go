package redisx

import (
	"context"
	"time"

	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// PipeWrite marshal objs to json string one by one, then using pipeline set to redis
func (c *Client) PipeWrite(ctx context.Context, objs []interface{}, kfn KeyFunc, ttl time.Duration) error {
	// make pipeline
	pipe := c.client.Pipeline()
	for _, obj := range objs {
		// serialize obj val
		val, err := jsoniter.MarshalToString(obj)
		if err != nil {
			continue
		}
		// set val
		pipe.Set(ctx, kfn(obj), val, ttl)
	}
	// exec pipeline
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "PipeWrite() pipeline exec error")
	}
	return nil
}

// PipeRead using pipeline read multi keys to map[string]string for each element, map key is redis key, map element value is redis value
func (c *Client) PipeRead(ctx context.Context, keys []string) (map[string]string, error) {
	if len(keys) == 0 {
		return nil, errors.New("PipeRead() read empty keys")
	}
	// make pipeline
	pipe := c.client.Pipeline()
	mcmd := map[string]*redis.StringCmd{}
	for _, key := range keys {
		mcmd[key] = pipe.Get(ctx, key)
	}
	// exec pipeline
	_, err := pipe.Exec(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.Wrap(err, "PipeRead() pipeline exec error")
	}

	// read from pipeline
	out := map[string]string{}
	for k, v := range mcmd {
		if err := v.Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			return nil, err
		}
		out[k] = v.Val()
	}

	return out, nil
}

// DelValues using redis client del multi keys
func (c *Client) DelValues(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	// 重试删除，避免失败
	var err error
	for i := 0; i < 3; i++ {
		if err = c.client.Del(ctx, keys...).Err(); err == nil { // 调用无错直接返回
			return nil
		}
		time.Sleep(time.Duration(i) * time.Millisecond)
	}
	return errors.Errorf("execution DelValues(%v) fails and retry reaches the max, %s", keys, err)
}
