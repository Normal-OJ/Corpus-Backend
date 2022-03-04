package kideval

import (
	"errors"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"main.main/src/db"
	"main.main/src/utils"
	"main.main/src/zhoseg"
)

func execute(speakers []string, files []string) (string, string, error) {
	cmdFolderLoc := os.Getenv("CLANG_CMD_FOLDER")
	chaCache := os.Getenv("CHA_CACHE")

	cmdOpts := []string{"+lzho"}
	for _, speaker := range speakers {
		cmdOpts = append(cmdOpts, "+t*"+speaker)
	}
	for _, file := range files {
		file = filepath.Clean(file)

		if !utils.PathChecker(file) && !utils.ChaCachePathChecker(file) {
			return "", "", errors.New("unallowed path")
		}
		cmdOpts = append(cmdOpts, file)
	}

	utils.RunCmd(cmdFolderLoc+"/kideval", cmdOpts)
	dat, err := ioutil.ReadFile(files[0][:len(files[0])-3] + "kideval.xls")
	if err != nil {
		return "", "", err
	}
	var out = string(dat)
	if !strings.Contains(out, "<?xml") {
		return "", "", errors.New(out)
	}

	file := strings.Split(out, "<?xml")[1]
	file = "<?xml" + strings.Split(file, "</Workbook>")[0] + "</Workbook>"

	filename := "kideval" + uuid.NewV4().String() + ".xls"
	ioutil.WriteFile(chaCache+"/"+filename, []byte(file), 0644)

	return filename, file, nil
}

func makeRespone(filename string, file string, indicator []string) map[string][]interface{} {
	data := utils.ExtractXMLInfo([]byte(file))
	ret := make(map[string][]interface{})
	ret["filename"] = []interface{}{filename}

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

	for k, v := range ret {
		if k == "filename" {
			continue
		}
		mean, n := utils.Mean(v)
		sd, _ := utils.SD(v)

		ret[k] = []interface{}{mean, sd, float64(n)}
	}

	return ret
}

func makeDetailedRespone(filename string, file string, chaFilename string) map[string]interface{} {
	data := utils.ExtractXMLInfo([]byte(file))
	ret := make(map[string]interface{})
	ret["filename"] = filename

	for _, row := range data[1:] {
		for index, val := range row {
			key := data[0][index].(string)
			ret[key] = val
		}
	}

	neededKeys := []string{"CTTR", "n_percentage", "v_percentage", "adj", "adj_percentage",
		"adv_percentage", "conj_percentage", "cl_percentage"}

	for _, key := range neededKeys {
		var val interface{}
		switch key {
		case "CTTR":
			val = ret["FREQ_types"].(float64) / math.Sqrt(ret["FREQ_tokens"].(float64)*2)
		case "adj":
			cmdFolderLoc := os.Getenv("CLANG_CMD_FOLDER")
			cmdOpts := []string{"+t%mor", "+sadj|*", chaFilename, "+t*CHI", "+d4"}

			out := utils.RunCmd(cmdFolderLoc+"/freq", cmdOpts)
			val, _ = strconv.Atoi(strings.Trim(strings.Split(out, "\n")[6][:5], " "))
		case "n_percentage":
			fallthrough
		case "v_percentage":
			fallthrough
		case "adj_percentage":
			fallthrough
		case "adv_percentage":
			fallthrough
		case "conj_percentage":
			fallthrough
		case "cl_percentage":
			word := strings.Split(key, "_")[0]
			val = utils.ToFloat(ret[word]) / utils.ToFloat(ret["mor_Words"])
		}
		ret[key] = val
	}

	return ret
}

type pathRequest struct {
	File      []string
	Speaker   []string
	Indicator []string
}

func getFiles(filename string) []string {
	finfo, _ := os.Stat(filename)
	ret := []string{}

	if finfo.IsDir() {
		files, _ := ioutil.ReadDir(filename)
		for _, file := range files {
			ret = append(ret, getFiles(filename+"/"+file.Name())...)
		}
	} else {
		if strings.HasSuffix(finfo.Name(), ".cha") {
			ret = append(ret, filename)
		}
	}

	return ret
}

// PathKidevalRequestHandler is like what it said :P
func PathKidevalRequestHandler(context *gin.Context) {
	var request pathRequest
	err := context.ShouldBind(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	files := []string{}

	for index, filename := range request.File {
		request.File[index] = utils.CHADIR + "/" + filename

		_, err := os.Stat(request.File[index])
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"message": "file: " + filename + " not found"})
			return
		}

		if !utils.PathChecker(filepath.Clean(request.File[index])) {
			context.JSON(http.StatusBadRequest, gin.H{"message": "unallowed path"})
			return
		}

		files = append(files, getFiles(request.File[index])...)
	}

	name, out, err := execute(request.Speaker, files)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ret := makeRespone(name, out, request.Indicator)

	context.JSON(http.StatusOK, ret)

}

type optionRequest struct {
	Ages      [][]int
	Sex       []int
	Context   []string
	Speaker   []string
	Indicator []string
}

// OptionKidevalRequestHandler is like what it said :P
func OptionKidevalRequestHandler(context *gin.Context) {
	var request optionRequest
	err := context.ShouldBind(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	var files = db.QueryChaFiles(request.Ages, request.Sex, request.Context)
	if len(files) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "filtered files' size is 0"})
		return
	}

	name, out, err := execute(request.Speaker, files)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ret := makeRespone(name, out, request.Indicator)

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
	err = context.ShouldBind(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	filename := "/tmp/" + uuid.NewV4().String() + ".cha"

	tmpFile, err := os.Create(filename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
		return
	}

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error})
		return
	}

	if !check(filename) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "File format error"})
		return
	}

	//mor have bug that can not support ???
	zhoseg.Call(filename)
	mor(filename)
	post(filename)
	postmortem(filename)
	megrasp(filename)
	name, out, err := execute(request.Speaker, []string{filename})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ret := makeRespone(name, out, request.Indicator)
	print(request.Indicator)
	print(request.Speaker)
	os.Remove(filename)

	context.JSON(http.StatusOK, ret)
}

type uploadDetailedRequest struct {
	Speaker []string
}

// UploadDetailedKidevalRequestHandler is like what it said :P
func UploadDetailedKidevalRequestHandler(context *gin.Context) {
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}

	var request uploadDetailedRequest
	err = context.ShouldBind(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	//NOTE: mor has bug that can not support any relative or absolute path
	filename := uuid.NewV4().String() + ".cha"
	tmpFile, err := os.Create(filename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
		return
	}
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error})
		return
	}
	if !check(filename) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "File format error"})
		return
	}
	zhoseg.Call(filename)
	mor(filename)
	post(filename)
	postmortem(filename)
	megrasp(filename)
	err = utils.MoveFile(filename, utils.CHACACHE+"/"+filename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	filename = utils.CHACACHE + "/" + filename
	name, out, err := execute(request.Speaker, []string{filename})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ret := makeDetailedRespone(name, out, filename)
	print(request.Speaker)
	os.Remove(filename)
	context.JSON(http.StatusOK, ret)
}
