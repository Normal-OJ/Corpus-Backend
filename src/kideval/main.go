package kideval

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"main.main/src/db"
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

func makeRespone(file string, indicator []string) map[string][]interface{} {
	data := utils.ExtractXMLInfo([]byte(file))
	ret := make(map[string][]interface{})

	for _, key := range indicator {
		ret[key] = make([]interface{}, 0)
	}

	for _, row := range data[1:] {
		for index, val := range row {
			key := data[0][index].(string)
			_, ok := ret[key]
			if ok {
				ret[key] = append(ret[key], val)
			}
		}
	}

	return ret
}

type pathRequest struct {
	File      []string
	Speaker   []string
	Indicator []string
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
	ret := makeRespone(out, request.Indicator)

	context.JSON(http.StatusOK, ret)

}

type optionRequest struct {
	Age       [][]int
	Sex       []int
	Context   []string
	Speaker   []string
	Indicator []string
}

// OptionKidevalRequestHandler is like what it said :P
func OptionKidevalRequestHandler(context *gin.Context) {
	var request optionRequest
	context.ShouldBind(&request)

	var files = db.QueryChaFiles(request.Age, request.Sex, request.Context)
	out := execute(request.Speaker, files)
	ret := makeRespone(out, request.Indicator)

	context.JSON(http.StatusOK, ret)
}

type uploadRequest struct {
	Speaker   []string
	Indicator []string
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

	tmpFile, err := os.Create("/tmp/a.xls")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error})
		return
	}

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error})
		return
	}

	out := execute(request.Speaker, []string{"/tmp/a.xls"})
	ret := makeRespone(out, request.Indicator)
	print(request.Indicator)
	print(request.Speaker)

	context.JSON(http.StatusOK, ret)
}
