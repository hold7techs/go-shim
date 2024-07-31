package retryx

import (
	"context"
	"fmt"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/hold7techs/go-shim/shim/log"
)

// RetryOptErr 重试操作错误
type RetryOptErr struct {
	err error
}

func (e *RetryOptErr) Error() string {
	return fmt.Sprintf("retry opt err: %v", e.err)
}

// RetryOptErrWrap 返回一个可重试操作的错误，用于重试检测
func RetryOptErrWrap(err error) *RetryOptErr {
	return &RetryOptErr{err: err}
}

// DoRetryFunc 依赖retry-go，失败即进行重试
func DoRetryFunc(ctx context.Context, retryableFunc func() error, retryTimes int, retryIfFn func(error) bool) error {
	return retry.Do(
		retryableFunc,
		retry.Attempts(uint(retryTimes)),
		retry.Delay(100*time.Millisecond),
		retry.MaxDelay(250*time.Millisecond),
		retry.LastErrorOnly(true),
		retry.RetryIf(retryIfFn),
		retry.OnRetry(func(n uint, err error) {
			log.Errorf("retry[#%d] because got err: %s", n, err)
		}),
	)
}
