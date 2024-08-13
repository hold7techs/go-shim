package shim

import (
	"github.com/lupguo/go-shim/x/log"
	"github.com/pkg/errors"
)

// LogAndWrapErr 打印日志并返回wrap的错误码信息
func LogAndWrapErr(err error, format string, args ...any) error {
	e := errors.Wrapf(err, format, args)
	log.Error(e)

	return e
}
