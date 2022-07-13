package utils

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
)

var node *snowflake.Node

func init() {
	var st time.Time
	st, err := time.Parse("2006-01-02", "2022-01-01")
	Panic(err)
	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(1)
	Panic(err)
	logrus.Info("init snowflake")
}

func GenID() int64 {
	return node.Generate().Int64()
}
