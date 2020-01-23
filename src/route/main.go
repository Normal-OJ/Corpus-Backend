package route

import (
	"github.com/gin-gonic/gin"
	"main.main/src/auth"
)

func RegisterRouter(engine *gin.Engine) {

	// register authenticated required funcs
	authGroup := engine.Group("/mod")
	authMid := auth.GetAuthMiddleware()

	authGroup.Use(authMid.MiddlewareFunc())
	{

	}
	engine.POST("login", authMid.LoginHandler)
}
