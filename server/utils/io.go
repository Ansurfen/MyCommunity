package utils

import (
	"github.com/spf13/viper"
)

var Conf *viper.Viper

func init() {
	Conf = getConf("application", "yaml", ".")
}

func getConf(confName, confType, dir string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigName(confName)
	conf.SetConfigType(confType)
	conf.AddConfigPath(dir)
	Panic(conf.ReadInConfig())
	return conf
}
