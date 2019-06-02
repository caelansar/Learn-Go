package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"webapp/config"
	"webapp/db"
	"webapp/model"
)

var (
	DbEngin *gorm.DB
)

func init() {
	var (
		err error
	)
	DbEngin, err = db.InitMysql(config.Cfg.MysqlUn, config.Cfg.MysqlPwd,
		config.Cfg.MysqlHost, config.Cfg.MysqlDb, config.Cfg.MysqlPort)
	if err != nil{
		fmt.Println(err)
		return
	}
	DbEngin.DropTable(model.User{}).AutoMigrate(model.User{})
	fmt.Println("init database success")
}
