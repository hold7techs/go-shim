package go_shim

import (
	"github.com/shopspring/decimal"
)

// ToYuan 分转成元浮点型
func ToYuan(o int64) float64 {
	d := decimal.New(1, 2) // 分除以100得到元
	res, _ := decimal.NewFromInt(o).DivRound(d, 2).Float64()
	return res
}

// ToFen 处理金额 元 转分
func ToFen(amount float64) int64 {
	d := decimal.New(1, 2)
	d1 := decimal.New(1, 0)
	// 当乘以100后，仍然有小数位，取四舍五入法后，再取整数部分
	dff := decimal.NewFromFloat(amount).Mul(d).DivRound(d1, 0).IntPart()
	return dff
}

// ToIntYuan 分转成整型元
func ToIntYuan(y int64) int64 {
	return y / 100
}
