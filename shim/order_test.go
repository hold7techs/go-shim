package shim

import (
	"testing"
)

func TestGenerateOrderSN(t *testing.T) {
	// 示例生成订单编号
	for i := 0; i < 10; i++ {
		t.Logf(GenerateSN("AC_", 5))
	}
}
