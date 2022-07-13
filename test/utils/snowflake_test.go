package utils

import (
	"MyCommunity/utils"
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	fmt.Println(utils.GenID())
}
