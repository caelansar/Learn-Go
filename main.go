package main

import (
	"github.com/gin-gonic/gin"
	"webapp/controller"
	"webapp/middleware"
)

var buildstamp = ""
var goversion = ""

type H struct {
	name string
	age int
}

func main() {
	route := gin.Default()
	v1 := route.Group("/v1")
	{
		v1.GET("/test", controller.Test)
		v1.GET("/hello", controller.Hello)
		userRouter := v1.Group("/user")
		{
			userRouter.POST("/register", controller.UserRegister)
			userRouter.POST("/login", controller.UserLogin)
			userRouter.POST("/info", middleware.Auth, controller.UserInfo)
		}
	}
	route.Run()
	//args := os.Args
	//if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
	//	fmt.Printf("Go Version : %s\n", goversion)
	//	fmt.Printf("UTC Build Time : %s\n", buildstamp)
	//	return
	//}
	//fmt.Println("hello world")
}