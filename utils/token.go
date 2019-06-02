package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"webapp/config"
	"webapp/model"
)

type Claims struct {
	jwt.StandardClaims
	Uid uint `json:"uid"`
}

func GenerateToken(user model.User) (string, error) {

	myClaims := &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.ID,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenStr, err := tokenClaims.SignedString([]byte(config.Cfg.SecretKey))
	return tokenStr, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Cfg.SecretKey), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}