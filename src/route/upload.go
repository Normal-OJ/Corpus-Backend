package route

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Upload(context *gin.Context) {
	prog := context.Query("prog")
	_, fileheader, _ := context.Request.FormFile("file")
	multi, _ := strconv.ParseBool(context.Request.FormValue("multi"))
	opts := context.Request.FormValue("opts")

	folderName := "/tmp/Req"
	os.Mkdir(folderName, 0777)
	// save file
	filename := "data.cha"
	if multi == true {
		filename = "data.zip"
	}
	context.SaveUploadedFile(fileheader, folderName+"/"+filename)

	if multi == true {
		context.String(http.StatusBadRequest, "not implemented :P\n")
		return
	}

}
