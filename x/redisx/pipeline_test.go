package redisx

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// PhoneAcc 已绑手机号账号信息，主要用于资产账号查下手机号
type PhoneAcc struct {
	AssetUID     uint64 // 资产UID
	AssetUIDType uint32 // 资产UIDType
	UID          uint64 // 账号uid，即tinyID
	UIDType      uint32 // 账号类型uid_type，因为是查手机号，有结果的话uid_type=1012
	NationCode   string // 手机号国家编码 86
	PhoneNumber  string // 绑定手机号
}

type Student struct {
	Name  string
	Grade uint64
	Class uint64
	ID    uint64
}

var (
	// params
	ctx = context.Background()

	// key
	rkUidPhone = `uid_phone:%d`
	rkStudent  = `stu:g:%d:client:%d:id:%d`
)

func TestWritePhoneAcc(t *testing.T) {
	// miniredis
	setup()
	defer teardown()
	hand := New(rdsClient)

	// data
	accs := []*PhoneAcc{
		{AssetUID: 1, AssetUIDType: 0, UID: 0, UIDType: 0, NationCode: "", PhoneNumber: "185"},
		{AssetUID: 2, AssetUIDType: 1, UID: 0, UIDType: 0, NationCode: "", PhoneNumber: "133"},
		{AssetUID: 3, AssetUIDType: 0, UID: 0, UIDType: 0, NationCode: "", PhoneNumber: "121"},
		{AssetUID: 4, AssetUIDType: 0, UID: 0, UIDType: 0, NationCode: "", PhoneNumber: "100002"},
	}
	var inputs []interface{}
	for _, acc := range accs {
		inputs = append(inputs, acc)
	}
	// key
	var kfn KeyFunc = func(i interface{}) string {
		if v, ok := i.(*PhoneAcc); ok {
			return RdKey(rkUidPhone, v.AssetUID)
		}
		return ""
	}
	// write
	err := hand.PipeWrite(ctx, inputs, kfn, redis.KeepTTL)
	assert.Nil(t, err)
}

func TestWriteStudent(t *testing.T) {
	// miniredis
	setup()
	defer teardown()
	hand := New(rdsClient)

	// data
	stus := []*Student{
		{Name: "user01", Grade: 3, Class: 2, ID: 1},
		{Name: "user02", Grade: 3, Class: 2, ID: 2},
		{Name: "user03", Grade: 4, Class: 1, ID: 1},
	}
	var inputs []interface{}
	for _, s := range stus {
		inputs = append(inputs, s)
	}
	// key
	var kfn KeyFunc = func(i interface{}) string {
		v := i.(*Student)
		return RdKey(rkStudent, v.Grade, v.Class, v.ID)
	}
	err := hand.PipeWrite(ctx, inputs, kfn, redis.KeepTTL)
	assert.Nil(t, err)
}

func TestReadPhoneAcc(t *testing.T) {
	// miniredis
	setup()
	defer teardown()
	c := New(rdsClient)

	// keys
	ids := []uint64{1, 2, 3}
	var keys []string
	for _, id := range ids {
		keys = append(keys, RdKey(rkUidPhone, id))
	}
	// out
	out, err := c.PipeRead(ctx, keys)
	assert.Nil(t, err)

	// scan to objs
	accs := make(map[uint64]*PhoneAcc)
	for _, v := range out {
		o := &PhoneAcc{}
		if err := json.Unmarshal([]byte(v), o); err != nil {
			t.Logf("unmarshal fail:%s", err)
		}
		accs[o.AssetUID] = o
	}

	for k, v := range accs {
		t.Logf("assetUID=>%d, phoneAccs=>%+v", k, *v)
	}
}

func TestRdCache_PipeWrite(t *testing.T) {
	// miniredis
	setup()
	defer teardown()
	c := New(rdsClient)

	type args struct {
		ctx  context.Context
		objs []interface{}
		kfn  KeyFunc
		t    time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"c1", args{ctx, []interface{}{1, 3, 5}, func(i interface{}) string { return fmt.Sprintf("i:%d", i) }, redis.KeepTTL}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.PipeWrite(tt.args.ctx, tt.args.objs, tt.args.kfn, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("PipeWrite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRdCache_PipeRead(t *testing.T) {
	type args struct {
		ctx  context.Context
		keys []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"c1", args{ctx: ctx, keys: []string{"i:3", "i:1", "uid_phone:1"}}, false},
	}

	// miniredis
	setup()
	defer teardown()
	c := New(rdsClient)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.PipeRead(tt.args.ctx, tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("PipeRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got map:%+v", got)
		})
	}
}
