package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

// Check error
func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// RequestHandler is like what it said :P
func RequestHandler(context *gin.Context) {
	filename := context.Query("file")

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	fi, err := os.Stat(filename)
	if err != nil {
		context.String(http.StatusNotFound, "The file does not exist")
		return
	}

	if fi.IsDir() {
		files, err := ioutil.ReadDir(filename)
		Check(err)

		var dirCount = 0
		var fileCount = 0

		for _, f := range files {
			fi, err := os.Stat(filename + "/" + f.Name())
			Check(err)

			if fi.IsDir() {
				dirCount++
			} else {
				fileCount++
			}
		}

		var folderNames = make([]string, dirCount)
		var fileNames = make([]string, fileCount)

		dirCount = 0
		fileCount = 0

		for _, f := range files {
			fi, err := os.Stat(filename + "/" + f.Name())
			Check(err)

			if fi.IsDir() {
				folderNames[dirCount] = f.Name()
				dirCount++
			} else {
				fileNames[fileCount] = f.Name()
				fileCount++
			}
		}

		context.JSON(http.StatusOK, gin.H{"folders": folderNames, "files": fileNames})
	} else {
		var dat, err = ioutil.ReadFile(filename)
		Check(err)

		context.JSON(http.StatusOK, gin.H{"context": string(dat)})
	}
}
