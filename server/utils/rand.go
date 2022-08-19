package utils

import (
	"math/rand"
	"time"
)

func RandInt(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}

func RandValue(args ...any) string {
	res := ""
	for _, v := range args {
		res += ToString(v)
	}
	res += ToString(RandInt(1000))
	return MD5(res)
}
