package route

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.main/src/utils"
)

func Upload(context *gin.Context) {
	//prog := context.Query("prog")
	_, fileheader, _ := context.Request.FormFile("file")
	multi, _ := strconv.ParseBool(context.Request.FormValue("multi"))
	//opts := context.PostFormArray("opts")

	folderName := "/tmp/Req" + utils.CreateFolderID()
	os.Mkdir(folderName, 0777)
	// save file
	filename := "data.cha"
	if multi == true {
		filename = "data.zip"
	}
	context.SaveUploadedFile(fileheader, folderName+"/"+filename)

	if multi == true {
		// utils.Unzip(folderName+"/"+filename , folderName)
		context.String(http.StatusBadRequest, "not implemented :P\n")
	} else {
		//clan.Exec(context, prog, opts, []string{folderName + "/" + filename})
	}
	os.RemoveAll(folderName)
}
