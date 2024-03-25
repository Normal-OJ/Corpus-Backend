package maps

import (
	"io"
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"os/exec"
	"github.com/gin-gonic/gin"
	"main.main/src/utils"
	uuid "github.com/satori/go.uuid"
)
func TestRequestHandler(context *gin.Context){
	sample := context.Query("file")

	resultFilename := uuid.NewV4().String() + ".csv"
	resultFile, err := os.Create(resultFilename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
		return
	}
	defer resultFile.Close()

	srcFile, err := os.Open("src/maps/sample/"+sample+"_res.csv")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
		return
	}
	defer srcFile.Close()

	_, err = io.Copy(resultFile, srcFile);
	if  err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
		return
	}
	
	err = utils.MoveFile(resultFilename, utils.CHACACHE+"/"+resultFilename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	jsonFile, err := ioutil.ReadFile("src/maps/sample/"+sample+"_res.json")
	if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
			return
	}

	// We use a map of interfaces to hold the JSON content
	var result map[string]interface{}
	err = json.Unmarshal(jsonFile, &result)
	if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
			return
	}

	// Update the ExportFileDir value
	result["ExportFileDir"] = resultFilename

	context.JSON(http.StatusOK, result)
}

func UploadMapsRequestHandler(context *gin.Context) {
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
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
	// ext := filepath.Ext(header.Filename)
	oriFilename := uuid.NewV4().String() + "txt"
	tmpFile, err := os.Create(oriFilename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error})
		return
	}

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error})
		return
	}
	tmpFile.Close()

	var jsonOutput bytes.Buffer
	var stderr bytes.Buffer

	resultFilename := "maps" + uuid.NewV4().String() + ".csv"
	cmd := exec.Command("python3", "src/maps/test.py", oriFilename, resultFilename)
	cmd.Stdout = &jsonOutput
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}

	// Unmarshal the JSON output
	var tableJson map[string]interface{}
	if err := json.Unmarshal(jsonOutput.Bytes(), &tableJson); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": "error processing JSON output"})
		return
	}
	tableJson["ExportFileDir"] = resultFilename

	err = utils.MoveFile(oriFilename, utils.CHACACHE+"/"+oriFilename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	oriFilename = utils.CHACACHE + "/" + oriFilename

	err = utils.MoveFile(resultFilename, utils.CHACACHE+"/"+resultFilename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	resultFilename = utils.CHACACHE + "/" + resultFilename

	context.JSON(http.StatusOK, tableJson)
}