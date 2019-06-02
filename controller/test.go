package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"webapp/helper"
)

func Test(c *gin.Context) {
	// get raw bytes
	d, err := c.GetRawData()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(d)
	rep := helper.Ret{1, "success", nil}
	c.JSON(200, rep)
}


func Hello(c *gin.Context) {
	// get raw bytes
	d, err := c.GetRawData()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(d)
	rep := helper.Ret{1, "hello", nil}
	c.JSON(200, rep)
}