package route

import (
	"github.com/gin-gonic/gin"
	"main.main/src/view"
)

//RegisterRouter register all the required router
func RegisterRouter(engine *gin.Engine) {

	// register authenticated required funcs
	//not done yet :P

	//register view routers
	engine.GET("/api/view", view.RequestHandler)

	// clan router
	engine.POST("/clan/:prog", clanRequestHandler)
}
