package crond

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestTask_DistributeExec(t *testing.T) {
	// mock
	ctrl := gomock.NewController(t)
	locker := NewMockLocker(ctrl)
	locker.EXPECT().Lock(gomock.Any(), gomock.Any()).MaxTimes(1).Return(nil)
	locker.EXPECT().Lock(gomock.Any(), gomock.Any()).AnyTimes().Return(errors.New("get lock fail"))
	locker.EXPECT().UnLock(gomock.Any()).AnyTimes()

	type fields struct {
		Name        string
		Express     string
		Exec        func()
		dlocker     DistributeLocker
		dlockExpire time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"t1", fields{
			Name:    "task1",
			Express: "",
			Exec: func() {
				t.Logf("exec task once")
				time.Sleep(500 * time.Millisecond)
			},
			dlocker:     locker,
			dlockExpire: 3 * time.Second,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			task := &Task{
				Name:    tt.fields.Name,
				Express: tt.fields.Express,
				Exec:    tt.fields.Exec,
				Locker:  tt.fields.dlocker,
				TTL:     tt.fields.dlockExpire,
			}
			for i := 0; i < 10; i++ {
				task.Execute()
			}
		})
	}
}
