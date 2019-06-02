package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
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
	DbEngin, err = db.InitMysql("root", "6666", "localhost", "webapp", "3306")
	if err != nil{
		fmt.Println(err)
		return
	}
	DbEngin.DropTable(model.User{}).AutoMigrate(model.User{})
	fmt.Println("init database success")
}
