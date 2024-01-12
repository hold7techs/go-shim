package redisx

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var redisServer *miniredis.Miniredis
var rdsClient *redis.Client

func setup() {
	var err error
	redisServer, err = miniredis.Run()
	if err != nil {
		panic(err)
	}
	rdsClient = redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
}

func teardown() {
	redisServer.Close()
}

func TestRdsHand_DistrbLock(t *testing.T) {
	// miniredis
	setup()
	defer teardown()

	// go mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hand := New(rdsClient)

	ctx := context.Background()
	key := "hello"

	// defer unlock
	defer hand.DistrbUnLock(ctx, key)

	// try lock, got lock
	err := hand.DistrbLock(ctx, key, 10*time.Second)
	assert.Nil(t, err)

	// try lock again, got error
	err = hand.DistrbLock(ctx, key, 10*time.Second)
	assert.NotNil(t, err)

	// unlock
	hand.DistrbUnLock(ctx, key)
	err = hand.DistrbLock(ctx, key, 10*time.Second)
	assert.Nil(t, err)
}
