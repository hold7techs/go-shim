package shim

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// SNPrefixHead SN编号前缀
type SNPrefixHead string

const (
	SNPrefixOrder          SNPrefixHead = "OD" // 订单
	SNPrefixPayment        SNPrefixHead = "PY" // 付款单
	SNPrefixDeliveryNote   SNPrefixHead = "DN" // 发货单
	SNPrefixLogisticsOrder SNPrefixHead = "LO" // 物流单
	SNPrefixAcceptance     SNPrefixHead = "AC" // 验收单
	SNPrefixRefund         SNPrefixHead = "RF" // 退款单
)

// GenerateSN 生成有序随机SN，例如订单格式为 SN_{时间戳}{订单号长度}, length为时间戳下单随机长度
func GenerateSN(prefix SNPrefixHead, length int) string {
	// 使用当前时间戳作为前缀部分
	timestamp := time.Now().Unix()

	// 生成随机订单号部分
	orderID := generateRandomOrderID(length) // 订单号长度为8

	// 拼接成最终的 OrderSN
	orderSN := fmt.Sprintf("%s%d%s", prefix, timestamp, orderID)

	return orderSN
}

// generateRandomOrderID 生成随机订单号，支持A-Z, 0-9组合
func generateRandomOrderID(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	builder := strings.Builder{}
	builder.Grow(length)

	for i := 0; i < length; i++ {
		randomChar := charset[seededRand.Intn(len(charset))]
		builder.WriteByte(randomChar)
	}

	return builder.String()
}
