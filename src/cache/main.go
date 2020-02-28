package cache

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"main.main/src/utils"
)

func RequestHandler(context *gin.Context) {
	target := context.Query("file")
	target = filepath.Clean(utils.CHADIR + "/" + target)
	target, err := filepath.Abs(target)
	if err != nil {
		context.String(http.StatusBadRequest, "invaild path")
		return
	}
	if !utils.PathChecker(target) {
		context.String(http.StatusForbidden, "invaild path")
		return
	}
	tmpFid := uuid.NewV4()

	err = utils.Zip(target, utils.CHACACHE+"/tmp"+tmpFid.String()+".zip")
	if err != nil {
		context.String(http.StatusInternalServerError, "error while creating zip")
		return
	}
	context.JSON(http.StatusOK, gin.H{"path": "tmp" + tmpFid.String() + ".zip"})
}
