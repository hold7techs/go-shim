package go_shim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJsonString(t *testing.T) {
	type Fav struct {
		Name string
	}
	type user struct {
		Id   int
		Name string
		Fav  []*Fav
	}
	t.Logf(ToJsonString(user{1, "user1", []*Fav{{"sport"}}}, true))
	t.Logf(ToJsonString(&user{1, "user1", []*Fav{{"sport"}}}, false))
	t.Logf("%+v", &user{1, "user1", []*Fav{{"sport"}}})

	type args struct {
		v      interface{}
		pretty bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{user{1, "user1", []*Fav{{"sport"}}}, true}, "{\n  \"Id\": 1,\n  \"Name\": \"user1\",\n  \"Fav\": [\n    {\n      \"Name\": \"sport\"\n    }\n  ]\n}"},
		{"t2", args{user{1, "user1", []*Fav{{"sport"}}}, false}, `{"Id":1,"Name":"user1","Fav":[{"Name":"sport"}]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ToJsonString(tt.args.v, tt.args.pretty), "ToJsonString(%v, %v)", tt.args.v, tt.args.pretty)
		})
	}
}
