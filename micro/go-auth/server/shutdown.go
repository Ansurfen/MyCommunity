package server

import (
	"go-auth/db/nosql"
	"go-auth/db/sql"
	"go-auth/utils"
	"os"
	"os/signal"
)

func ShutDown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	defer nosql.GetRedis().Close()
	dbs := utils.Conf.GetStringSlice("datasource.database")
	for _, db := range dbs {
		sql.ClearDB(db)
	}
	os.Exit(0)
}
