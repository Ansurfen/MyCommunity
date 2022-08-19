package sql

import (
	"MyCommunity/utils"
	_ "github.com/go-sql-driver/mysql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type DB struct {
	*gorm.DB
}

func init() {
	dbs := utils.Conf.GetStringSlice("datasource.database")
	for _, db := range dbs {
		SetDB(db,initDB(getArgs(db)))
	}
}

func initDB(args string) *DB {
	db, err := gorm.Open(utils.Conf.GetString("datasource.driver"), args)
	utils.Panic(err)
	err = db.DB().Ping()
	utils.Panic(err)
	logrus.Info("Successfully connected!")
	return &DB{
		DB: db,
	}
}

func getArgs(database string) string {
	host := utils.Conf.GetString("datasource.host")
	port := utils.Conf.GetString("datasource.port")
	username := utils.Conf.GetString("datasource.username")
	password := utils.Conf.GetString("datasource.password")
	charset := utils.Conf.GetString("datasource.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
}
