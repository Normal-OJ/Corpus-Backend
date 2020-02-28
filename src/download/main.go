package download

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"main.main/src/utils"
)

func RequestHandler(context *gin.Context) {
	target := context.Query("file")
	target = filepath.Clean(utils.CHACACHE + "/" + target)
	target, err := filepath.Abs(target)
	if err != nil {
		context.String(http.StatusBadRequest, "invaild path")
		return
	}
	if !utils.ChaCachePathChecker(target) {
		context.String(http.StatusForbidden, "invaild path")
		return
	}
	finfo, err := os.Stat(target)
	if err != nil || finfo.IsDir() {
		context.String(http.StatusNotFound, "can not get file")
		return
	}
	context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(target))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	context.Writer.Header().Add("Content-Type", "application/octet-stream")
	context.File(target)
	context.String(http.StatusOK, "OK")
	return
}
