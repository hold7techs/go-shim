package grayacc

import (
	"testing"
)

func TestHitThresholdRule(t *testing.T) {
	type args struct {
		uid       uint64
		threshold uint64
		grayUIDs  []uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"wx_00_1", args{144115216610503700, 0, nil}, false},
		{"wx_00_2", args{144115216610503790, 0, nil}, false},
		{"wx_00_3", args{144115216610503799, 0, nil}, false},
		{"wx_2", args{144115216610503700, 1, nil}, true},
		{"wx_3", args{144115216610503700, 100, nil}, true},
		{"wx_4", args{144115216610503799, 100, nil}, true},
		{"wx_5", args{144115216610503799, 99, nil}, false},
		{"qq_00-09", args{287853405, 10, nil}, true},
		{"qq_00-05", args{287853405, 5, nil}, false},
		{"qq_gray_no_exist", args{287853405, 5, []uint64{1, 2, 3}}, false},
		{"qq_gray_exist_1", args{287853405, 5, []uint64{287853405, 2, 3}}, true},
		{"qq_gray_exist_2", args{287853405, 0, []uint64{1, 2, 287853405}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HitThreshold(tt.args.uid, tt.args.threshold, tt.args.grayUIDs); got != tt.want {
				t.Errorf("hitGrayRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
