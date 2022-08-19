package utils

import (
	"strconv"
	"time"
)

func NowTimestamp() int64 {
	return time.Now().Unix()
}

func NowTimestampByString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
