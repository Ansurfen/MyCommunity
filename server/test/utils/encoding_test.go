package utils

import (
	"MyCommunity/utils"
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	var i64 int64 = 123456
	var i32 int = 123456
	var f64 float64 = 100.123456789123456789
	fmt.Println(utils.ToString(i64), utils.ToString(i32), utils.ToString(f64))
}
