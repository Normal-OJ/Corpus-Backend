package route

import (
	"github.com/gin-gonic/gin"
	"main.main/src/cache"
	"main.main/src/download"
	"main.main/src/kideval"
	"main.main/src/modify"
	"main.main/src/view"
	"main.main/src/zhoseg"
)

//RegisterRouter register all the required router
func RegisterRouter(engine *gin.Engine) {

	// register authenticated required funcs
	//not done yet :P

	//register function routers
	engine.GET("/api/view", view.RequestHandler)
	engine.POST("/api/upload", modify.UploadRequestHandler)
	engine.POST("/api/option_kideval", kideval.OptionKidevalRequestHandler)
	engine.POST("/api/path_kideval", kideval.PathKidevalRequestHandler)
	engine.POST("/api/upload_kideval", kideval.UploadKidevalRequestHandler)
	engine.POST("/api/upload_detailed_kideval", kideval.UploadDetailedKidevalRequestHandler)

	//register download routers
	engine.GET("/api/download", download.RequestHandler)

	//register zipping routers
	engine.POST("/api/zip", cache.RequestHandler)

	engine.POST("/api/segment", zhoseg.UploadSegmentHandler)
}
