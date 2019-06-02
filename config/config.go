package config

import "github.com/Terry-Mao/goconf"

var Cfg Config

type Config struct {
	SecretKey string `goconf:"secret:key"`
	MysqlHost string `goconf:"mysql:host" split_words:"false"`
	MysqlUn   string `goconf:"mysql:username" split_words:"false"`
	MysqlPwd  string `goconf:"mysql:password" split_words:"false"`
	MysqlPort string `goconf:"mysql:port" split_words:"false"`
	MysqlDb   string `goconf:"mysql:db"`
}

func init() {
	conf := goconf.New()
	if err := conf.Parse(".conf"); err != nil {
		panic(err)
	}
	err := conf.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
}