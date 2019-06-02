package db

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMysql(user, password, host, dbname, port string) (*gorm.DB, error) {
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err.Error())
		return db, err
	}
	return db, nil
}
