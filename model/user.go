package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Pwd  string `json:"pwd"`
	Salt string `json:"salt"`
}

