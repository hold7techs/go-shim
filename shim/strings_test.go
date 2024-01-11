package shim

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessStringsSlice(t *testing.T) {
	type args struct {
		strs   []string
		filter func(string) bool
		fn     func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{
			strs: []string{"1", "2"},
		}, []string{"1", "2"}},
		{"t2", args{
			strs: []string{"1", "2"},
			filter: func(s string) bool {
				return s == "1"
			},
			fn: nil,
		}, []string{"2"}},
		{"t3", args{
			strs: []string{"1", "2"},
			fn: func(s string) string {
				return fmt.Sprintf("0x%v", s)
			},
		}, []string{"0x1", "0x2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ProcessStringsSlice(tt.args.strs, tt.args.filter, tt.args.fn), "ProcessStringsSlice(%v, %v, %v)", tt.args.strs, tt.args.filter, tt.args.fn)
		})
	}
}
