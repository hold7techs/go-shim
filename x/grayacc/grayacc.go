// Package grayacc TODO
package grayacc

// HitThreshold 灰度规则控制
func HitThreshold(uid uint64, threshold uint64, grayUIDs []uint64) bool {
	// 尾号限制
	if uid%100 < threshold {
		return true
	}

	// 白名单限制
	for _, grayUID := range grayUIDs {
		if uid == grayUID {
			return true
		}
	}
	return false
}

// HitWhiteList 是否命中白名单
func HitWhiteList(uid uint64, uidList []uint64) bool {
	for _, u := range uidList {
		if uid == u {
			return true
		}
	}
	return false
}

// IsHitRules 是否命中了规则
func IsHitRules(rules []func() bool) bool {
	for _, rule := range rules {
		if rule() {
			return true
		}
	}
	return false
}
