package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.main/src/utils"
	"main.main/src/view"
	"net/http"
	"os"
	"strconv"
	"time"
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
		cmdFolderLoc = "/home/asef18766/桌面/LanguageDB/BackEnd/unix-clan/unix/bin"
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

	folderName := "/tmp/Req" + time.Now().Format("20060102150405")
	os.Mkdir(folderName, 0777)
	// save file
	filename := "data.cha"
	if multi == true {
		filename = "data.zip"
	}
	context.SaveUploadedFile(fileheader, folderName+"/"+filename)
	var cmdOpts []string
	if opts != "" {
		cmdOpts = append(cmdOpts, opts)
	}
	if multi {
		print("into multi\n")
		utils.Unzip(folderName+"/"+filename, folderName)
		cmdOpts = append(cmdOpts, "*.cha")
	} else {
		cmdOpts = append(cmdOpts, folderName+"/"+filename)
	}
	print("argc:", len(cmdOpts), "\n")
	fmt.Printf("exec command: %s %v\n", cmdFolderLoc+"/mlt", cmdOpts)

	var output = utils.RunCmd(cmdFolderLoc+"/mlt", cmdOpts)
	//os.RemoveAll(folderName)
	context.String(http.StatusOK, string(output))
	return
}
