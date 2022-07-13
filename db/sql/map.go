package sql

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var dmap sync.Map

func GetDB(database string) *DB {
	v, ok := dmap.Load(database)
	if !ok {
		logrus.Error("Fail to search db.")
	}
	return v.(*DB)
}

func ClearDB(database string) {
	v, ok := dmap.Load(database)
	if !ok {
		logrus.Error("Fail to search db to release.")
	}
	defer v.(*DB).Close()
	dmap.Delete(database)
}

func SetDB(database string, db *DB) {
	dmap.LoadOrStore(database, db)
}
