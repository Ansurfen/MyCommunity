package cache

import (
	"MyCommunity/cache"
	"MyCommunity/models"
	"MyCommunity/utils"
	"fmt"
	"runtime/debug"
	"testing"
)

func TestCache(t *testing.T) {
	debug.SetGCPercent(20)
	key := []byte("abc")
	val := []byte("def")
	cache.Cache.Set(key, val, cache.Min)
	got, err := cache.Cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", got)
	}
	affected := cache.Cache.Del(key)
	fmt.Println("deleted key ", affected)
	fmt.Println("entry count ", cache.Cache.EntryCount())
}

func TestStruct(t *testing.T) {
	key := utils.GenID()
	user := models.User{
		Username:  "root",
		Password:  "root",
		Telephone: "123456789",
		Id:        key,
	}
	val := utils.Serialize(user)
	cache.Cache.Set(utils.ToBytes(key), val, cache.Hour)
	tmpl := models.User{}
	got, _ := cache.Cache.Get(utils.ToBytes(key))
	utils.Unserialize(string(got), &tmpl)
	fmt.Println(tmpl)
}
