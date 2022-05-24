package mbase

import "math/rand"

// 随机值的返回 [min, max)
func RandInt64(min, max int64) int64 {
	if min >= max {
		return max
	}
	return rand.Int63n(max-min) + min
}
