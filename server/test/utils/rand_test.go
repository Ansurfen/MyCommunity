package utils

import (
	"MyCommunity/utils"
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	fmt.Println(utils.RandInt(100))
	fmt.Println(utils.RandValue("164asdafa45+95",1517545))
}
