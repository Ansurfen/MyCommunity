package nosql

import (
	"go-auth/utils"

	"github.com/go-redis/redis"
)

type Redis struct {
	*redis.Client
}

var redisdb Redis

func init() {
	host := utils.Conf.GetString("redis.host")
	port := utils.Conf.GetString("redis.port")
	password := utils.Conf.GetString("redis.password")
	redisdb.Client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
}

func GetRedis() Redis {
	return redisdb
}
