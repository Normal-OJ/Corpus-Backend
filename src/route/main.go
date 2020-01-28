package route

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"main.main/src/utils"
	"main.main/src/view"
)

//RegisterRouter register all the required router
func RegisterRouter(engine *gin.Engine) {

	// register authenticated required funcs
	//not done yet :P

	//register function routers
	engine.POST("/api/mlt", MltRequestHandler)

	//register view routers
	engine.GET("/api/view", view.RequestHandler)
}

// MltRequestHandler is like what it said :P
func MltRequestHandler(context *gin.Context) {
	cmdFolderLoc := os.Getenv("CLANG_CMD_FOLDER")

	if cmdFolderLoc == "" {
		// someone's home dir :P
		cmdFolderLoc = "/Users/chenzhangling/Desktop/unix-clan/unix/bin"
	}

	print("cmdFolderLoc:", cmdFolderLoc, "\n")
	/*defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusBadRequest, "request body missing")
			return
		}
	}()*/
	_, fileheader, _ := context.Request.FormFile("data")
	opts := context.Request.FormValue("opts")
	multi, _ := strconv.ParseBool(context.Request.FormValue("multi"))
	excel, _ := strconv.ParseBool(context.Request.FormValue("excel"))
	if excel == true || multi == true {
		context.String(http.StatusBadRequest, "not implemented :P\n")
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			fmt.Println("Recovered in f", err)
			return
		}
	}()

	folderName := "/Users/chenzhangling/Desktop/languageDB/BackEnd/Req" + time.Now().Format("20060102150405")
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
