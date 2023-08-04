package shim

import (
	"testing"
)

func TestToYuan(t *testing.T) {
	type args struct {
		o int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"t1", args{100}, 1},
		{"t2", args{10}, 0.1},
		{"t2", args{1}, 0.01},
		{"t4", args{105}, 1.05},
		{"t5", args{104}, 1.04},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToYuan(tt.args.o); got != tt.want {
				t.Errorf("ToYuan() = %v, want %v", got, tt.want)
			}
		})
	}
}
