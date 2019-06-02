package controller

import (
	"github.com/gin-gonic/gin"
	"webapp/helper"
	"webapp/service"
)

var (
	userService = service.UserService{}
)

func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	plainpwd := c.PostForm("password")
	user, err := userService.Register(username, plainpwd)
	if err != nil {
		rep := helper.Ret{Code: 1, Msg: err.Error()}
		c.JSON(200, rep)
		return
	}
	rep := helper.Ret{Code: 1, Msg: "success", Data:user}
	c.JSON(200, rep)
}