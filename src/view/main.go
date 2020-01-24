package view

import (
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"main.main/src/utils"
)

// RequestHandler is like what it said :P
func RequestHandler(context *gin.Context) {
	filename := context.Param("name")

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	if ioutil.IsDir(filename) {
		files, err := ioutil.ReadDir(filename)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
				fmt.Println(f.Name())
		}
	} else {

	}

	folderName := "/tmp/Req" + time.Now().Format("20060102150405")
	os.Mkdir(folderName, 0777)
	// save file
	filename := "data.cha"
	if multi == true {
		filename = "data.zip"
	}
	context.SaveUploadedFile(fileheader, folderName+"/"+filename)
	var cmd *exec.Cmd
	if multi == true {
		//  extract then execute(WIP)
		utils.Unzip(folderName+filename, folderName)
		cmd = exec.Command(cmdFolderLoc+"/mlt", opts, "*.cha")
	} else {
		cmd = exec.Command(cmdFolderLoc+"/mlt", opts, folderName+"/"+filename)
	}
	output, err := cmd.Output()
	if err != nil {
		print(err.Error())
		context.String(http.StatusInternalServerError, "command error")
		return
	}
	os.RemoveAll(folderName)
	context.String(http.StatusOK, string(output))
	return
}
