package shim

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	type args[T comparable] struct {
		ids []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{"t1", args[int]{[]int{1, 2, 1}}, []int{1, 2}},
		{"t2", args[int]{[]int{1, 3, 2, 1}}, []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uniq(tt.args.ids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInSlice(t *testing.T) {
	type args[T comparable] struct {
		id  T
		ids []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{"t1", args[int]{1, []int{1, 2}}, true},
		{"t2", args[int]{0, []int{1, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.id, tt.args.ids); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageData(t *testing.T) {
	type args[T interface{}] struct {
		data     []T
		page     int
		pageSize int
	}
	type testCase[T interface{}] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{"t1", args[int]{[]int{1, 2, 3, 4}, 1, 2}, []int{1, 2}},
		{"t2", args[int]{[]int{1, 2, 3, 4}, 2, 2}, []int{3, 4}},
		{"t3", args[int]{[]int{1, 2, 3, 4}, 2, 3}, []int{4}},
		{"t4", args[int]{[]int{1, 2, 3, 4}, 0, 2}, []int{1, 2}},
		{"t5", args[int]{[]int{1, 2, 3, 4}, 1, 10}, []int{1, 2, 3, 4}},
		{"t6", args[int]{[]int{1, 2, 3, 4}, 2, 10}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageData(tt.args.data, tt.args.page, tt.args.pageSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchNumbers(t *testing.T) {
	type args[T comparable] struct {
		ids  []T
		size int
	}
	type testCase[T comparable] struct {
		name        string
		args        args[T]
		wantBatches [][]T
	}
	tests := []testCase[int]{
		{name: "t1", args: args[int]{[]int{1, 2, 3}, 1}, wantBatches: [][]int{{1}, {2}, {3}}},
		{name: "t2", args: args[int]{[]int{1, 2, 3}, 2}, wantBatches: [][]int{{1, 2}, {3}}},
		{name: "t3", args: args[int]{[]int{1, 2, 3}, 3}, wantBatches: [][]int{{1, 2, 3}}},
		{name: "t4", args: args[int]{[]int{1, 2, 3}, 4}, wantBatches: [][]int{{1, 2, 3}}},
		{name: "t5", args: args[int]{[]int{1, 2, 3}, -2}, wantBatches: [][]int{{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantBatches, ShardingNumbers(tt.args.ids, tt.args.size), "ShardingNumbers(%v, %v)", tt.args.ids, tt.args.size)
		})
	}
}

func TestJoinNumbers(t *testing.T) {
	type args[T comparable] struct {
		ids []T
		sep string
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		{"t1", args[int]{[]int{1, 2, 3}, ","}, "1,2,3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, JoinNumbers(tt.args.ids, tt.args.sep), "JoinNumbers(%v, %v)", tt.args.ids, tt.args.sep)
		})
	}
}
