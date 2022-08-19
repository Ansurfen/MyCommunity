package server

import (
	"MyCommunity/cache"
	"MyCommunity/db/nosql"
	"MyCommunity/db/sql"
	"MyCommunity/utils"
	"os"
	"os/signal"
)

func ShutDown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	defer cache.Cache.Clear()
	defer nosql.GetRedis().Close()
	dbs := utils.Conf.GetStringSlice("datasource.database")
	for _, db := range dbs {
		sql.ClearDB(db)
	}
	os.Exit(0)
}
