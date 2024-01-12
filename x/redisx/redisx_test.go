package redisx

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestSimpleSet(t *testing.T) {
	setup()
	defer teardown()

	rds := rdsClient
	key := RdKey(`simple:%d`, time.Now().Unix())
	_, err := rds.Get(ctx, key).Bool()
	assert.Equal(t, err, redis.Nil)

	// set and get
	rds.Set(ctx, key, true, 30*time.Minute)
	val, err := rds.Get(ctx, key).Bool()
	assert.NotEqual(t, err, redis.Nil)
	assert.Equal(t, val, true)

	// set again
	rds.Set(ctx, key, false, 30*time.Minute)
	val, err = rds.Get(ctx, key).Bool()
	assert.NotEqual(t, err, redis.Nil)
	assert.Equal(t, val, false)
}

func TestHashSet(t *testing.T) {
	// miniredis
	setup()
	defer teardown()
	hand := New(rdsClient)

	rds := hand.GetClient()
	key := RdKey(`hashexp:%d`, time.Now().Unix())

	// err := redisServer.HMSet(ctx, "map",
	// 	"name", "hello",
	// 	"count", 123,
	// 	"correct", true).Err()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Get the map. The same approach works for HmGet().
	// res := redisServer.HGetAll(ctx, "map")
	// if res.Err() != nil {
	// 	panic(err)
	// }
	//
	// type data struct {
	// 	Name    string `redis:"name"`
	// 	Count   int    `redis:"count"`
	// 	Correct bool   `redis:"correct"`
	// }
	//
	// // Scan the results into the struct.
	// var d data
	// if err := res.Scan(&d); err != nil {
	// 	panic(err)
	// }
	//
	// t.Log(d)

	// hgetall
	var us UserList
	ma := rds.HGetAll(ctx, key).Val()
	for k, v := range ma {
		t.Logf("k=>%v, v=>%+v", k, v)
		uv := &User{}
		err := json.Unmarshal([]byte(v), uv)
		assert.Nil(t, err)
		us = append(us, uv)
	}

	// hsetall
	// mus := []*User{
	// 	{78, "clark"},
	// 	{32, "terry"},
	// 	{90, "raft"},
	// }
	// u := &User{77, "gght"}
	// set := redisServer.HMSet(ctx, key, u)
	// assert.Nil(t, set.Err())

	// hgetall
	var usm UserList
	ma = rds.HGetAll(ctx, key).Val()
	for k, v := range ma {
		t.Logf("k=>%v, v=>%+v", k, v)
		uv := &User{}
		err := json.Unmarshal([]byte(v), uv)
		assert.Nil(t, err)
		usm = append(usm, uv)
	}
	t.Logf("usm:+%+v", usm)
}

type User struct {
	ID   uint64 `redis:"id"`
	Name string `redis:"name"`
}

type UserList []*User

func (u *UserList) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}
