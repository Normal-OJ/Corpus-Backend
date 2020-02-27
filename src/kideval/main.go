package kideval

import (
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"main.main/src/db"
	"main.main/src/modify"
	"main.main/src/utils"
)

func execute(speakers []string, files []string) string {
	cmdFolderLoc := os.Getenv("CLANG_CMD_FOLDER")

	if cmdFolderLoc == "" {
		// someone's home dir :P
		if runtime.GOOS == "darwin" {
			cmdFolderLoc = "/Users/chenzhangling/Desktop/unix-clan/unix/bin"
		} else {
			cmdFolderLoc = "/home/asef18766/桌面/LanguageDB/BackEnd/unix-clan/unix/bin"
		}
	}

	cmdOpts := []string{"+lzho"}
	for _, speaker := range speakers {
		cmdOpts = append(cmdOpts, "+t*"+speaker)
	}
	for _, file := range files {
		cmdOpts = append(cmdOpts, file)
	}

	var out = utils.RunCmd(cmdFolderLoc+"/kideval", cmdOpts)
	file := strings.Split(out, "<?xml")[1]
	file = "<?xml" + strings.Split(file, "</Workbook>")[0] + "</Workbook>"

	id := utils.CreateID(file)
	ioutil.WriteFile("/tmp/kideval"+strconv.FormatInt(id, 10)+".xls", []byte(file), 0644)

	return file
}

type pathRequest struct {
	File      []string
	Speaker   []string
	Indicator []bool
}

// PathKidevalRequestHandler is like what it said :P
func PathKidevalRequestHandler(context *gin.Context) {
	var request pathRequest
	context.ShouldBind(&request)

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	out := execute(request.Speaker, request.File)

	context.Writer.WriteHeader(http.StatusOK)
	context.Header("Content-Disposition", "attachment; filename=kideval.xls")
	context.Writer.Write([]byte(out))

	/*
		context.Header("Content-Type", "multipart/mixed; boundary='@@@'")
		context.File("data.cha")*/
	//context.Writer.Write([]byte("@@@"))

}

type optionRequest struct {
	Age       [][]int
	Sex       []int
	Context   []string
	Speaker   []string
	Indicator []bool
}

// OptionKidevalRequestHandler is like what it said :P
func OptionKidevalRequestHandler(context *gin.Context) {
	var request optionRequest
	context.ShouldBind(&request)

	var files = db.QueryChaFiles(request.Age, request.Sex, request.Context)
	out := execute(request.Speaker, files)

	context.JSON(http.StatusOK, gin.H{"result": out})
}

type uploadRequest struct {
	Speaker []string
}

// UploadKidevalRequestHandler is like what it said :P
func UploadKidevalRequestHandler(context *gin.Context) {
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}

	var request uploadRequest
	context.ShouldBind(&request)

	err = modify.Upload(file, "data.cha")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error})
		return
	}

	out := execute(request.Speaker, []string{"data.cha"})

	context.JSON(http.StatusOK, gin.H{"result": out})
}
