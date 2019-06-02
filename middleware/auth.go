package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"webapp/helper"
	"webapp/utils"
)

func Auth(c *gin.Context) {
	// validate token
	//token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
	//	func(token *jwt.Token) (i interface{}, e error) {
	//		return []byte(config.Cfg.SecretKey), nil
	//})
	//if err == nil{
	//	err := token.Claims.Valid()
	//	if err == nil {
	//		c.Next()
	//	}else {
	//		c.Abort()
	//		rep := helper.Ret{-1, "token is not validate: " + err.Error(), nil}
	//		c.JSON(http.StatusOK, rep)
	//		return
	//	}
	//} else {
	//	c.Abort()
	//	rep := helper.Ret{-1, "access token is missing", nil}
	//	c.JSON(http.StatusOK, rep)
	//	return
	//}
	auth := c.Request.Header.Get("Authorization")
	token := strings.Split(auth, " ")[1]

	if token == "" {
		c.Abort()
		rep := helper.Ret{-1, "access token is missing", nil}
		c.JSON(http.StatusOK, rep)
		return
	} else {
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.Abort()
			rep := helper.Ret{-1, "token is invalid", nil}
			c.JSON(http.StatusOK, rep)
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.Abort()
			rep := helper.Ret{-1, "token is expired", nil}
			c.JSON(http.StatusOK, rep)
			return
		}
		c.Next()
	}

}