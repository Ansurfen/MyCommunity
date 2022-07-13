package db

import (
	"MyCommunity/db/nosql"
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	_, err := nosql.GetRedis().Client.Ping().Result()
	if err != nil {
		fmt.Printf("connect redis failed! err : %v\n", err)
		return
	}
	fmt.Println("connect redis successfully!")
}
