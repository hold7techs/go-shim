package shim

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTimeVersion(t *testing.T) {
	t.Logf(GetTimeVersion())
}

func TestStdDateTimeStr_GetTime(t *testing.T) {
	tests := []struct {
		name string
		s    StdDateTimeStr
		want time.Time
	}{
		{"t1", "2023-01-01 00:00:01", time.Date(2023, 1, 1, 0, 0, 1, 0, time.Local)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
