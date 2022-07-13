package common

import (
	"MyCommunity/common"
	"fmt"
	"testing"
)

func TestCheckPhone(t *testing.T) {
	fmt.Println(common.CheckPhone(`15000000000`))
	fmt.Println(common.CheckPhone(`11100000000`))
}

func TestCheckEmail(t *testing.T){
	fmt.Println(common.CheckEmail(`ansurfen@gmail`))
	fmt.Println(common.CheckEmail(`ansurfen@gmail.com`))
}