package mysqlx

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// IsRecordNotFoundErr 是否
func IsRecordNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
