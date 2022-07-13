package cache

import (
	"MyCommunity/utils"
	"strings"

	"github.com/coocood/freecache"
	"github.com/sirupsen/logrus"
)

var Cache *freecache.Cache

func init() {
	Cache = freecache.NewCache(getCacheSize())
	logrus.Info("init cache")
}

func getCacheSize() int {
	cacheSize := utils.Conf.GetFloat64("cache.size")
	switch _uint := utils.Conf.GetString("cache.uint"); strings.ToLower(_uint) {
	case "byte":
		cacheSize *= _byte
	case "b":
		cacheSize *= b
	case "kb":
		cacheSize *= kb
	case "mb":
		cacheSize *= mb
	case "gb":
		cacheSize *= gb
	default:
	}
	return int(cacheSize)
}
