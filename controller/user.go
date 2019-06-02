package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusOK, rep)
		return
	}
	rep := helper.Ret{Code: 1, Msg: "success", Data:user}
	c.JSON(http.StatusOK, rep)
}

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	plainpwd := c.PostForm("password")
	tokenStr, err := userService.Login(username, plainpwd)
	if err != nil {
		rep := helper.Ret{-1, err.Error(), nil}
		c.JSON(http.StatusOK, rep)
		return
	}
	data := map[string]string{
		"token": tokenStr,
	}
	rep := helper.Ret{1, "login success", data}
	c.JSON(http.StatusOK, rep)
}


func UserInfo(c *gin.Context) {
	rep := helper.Ret{1, "protected resource", nil}
	c.JSON(http.StatusOK, rep)
}