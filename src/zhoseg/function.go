package zhoseg

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"net/http"

	"github.com/gin-gonic/gin"
	"main.main/src/utils"
)

var zhosegFilePath string = os.Getenv("ZHOSEG_FILE_PATH")
var segment = "seg"
var mode = "tt"

func Call(inputFile string) {
	input := filepath.Base(inputFile) //extract only the input file name from the absolute direction
	outputDir, err := os.MkdirTemp("", "ex-zhoseg")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.RemoveAll(outputDir)
	temp, err := os.MkdirTemp("", "temp-zhoseg")
	if err != nil {
		log.Fatal(err)
		return
	}
	path_ := filepath.Join(temp, input) //generate a temparary folder direction
	fin, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fin.Close()
	fout, err := os.Create(path_)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fout.Close()
	_, err = io.Copy(fout, fin) //copy inputfile to the temp directiry
	if err != nil {
		log.Fatal(err)
		return
	}
	opt := []string{"zhoseg.js", mode, segment, temp, outputDir}
	//str := strings.Join(command, " ")

	parent := filepath.Dir(inputFile)                   //get the parent filepath
	parent = filepath.Base(parent)                      //extract the parent filename
	outputPath := filepath.Join(outputDir, temp, input) //add parent filename to filepath output filepath
	utils.RunCmdDir("node", opt, zhosegFilePath)
	//log.Fatal(opt)
	e := utils.MoveFile(outputPath, inputFile)
	if e != nil {
		log.Fatal(e)
	}

}

func UploadSegmentHandler(context *gin.Context) {
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}

	dirName, err := os.MkdirTemp("", "temp")
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer os.RemoveAll(dirName)

	filename := filepath.Join(dirName, "inputData.cha")
	out, err := os.Create(filename)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return
	}
	Call(filename)

	//context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base((filename))))
	//context.Writer.Header().Add("Content-Type", "plain/text; charset=\"utf-8\"")
	context.File(filename)
	//context.String(http.StatusOK, "OK")
}
