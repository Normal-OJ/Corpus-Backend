package auth

import (
	"crypto/rand"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	jwt2 "gopkg.in/dgrijalva/jwt-go.v3"
	"main.main/src/db"
	"time"
)

type login struct {
	username string
	passwd   string
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*login); ok {
		return jwt.MapClaims{
			"username": v.username,
			"passwd":   v.passwd,
		}
	}
	return jwt.MapClaims{}
}
func identityFunc(data jwt2.MapClaims) interface{} {
	return &login{
		username: data["username"].(string),
		passwd:   data["passwd"].(string),
	}
}
func authFunc(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginVals.username
	passwd := loginVals.passwd

	if _, err := db.FindUser(username, passwd); err != nil {
		return &login{
			username: username,
			passwd:   passwd,
		}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}
func checkFunc(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*login); ok {
		if _, err := db.FindUser(v.username, v.passwd); err != nil {
			return false
		}
		return true
	}
	return false
}
func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	key := make([]byte, 64)
	if i, e := rand.Read(key); e != nil {
		fmt.Printf("exception happened while creating key , create key %d bytes , error:%s\n", i, e.Error())
		return nil
	}
	return &jwt.GinJWTMiddleware{
		Realm:            "NTNULanguageDB",
		Key:              key,
		SigningAlgorithm: "HS512",
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		PayloadFunc:      payloadFunc,
		IdentityHandler:  identityFunc,
		Authenticator:    authFunc,
		Authorizator:     checkFunc,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	}
}
