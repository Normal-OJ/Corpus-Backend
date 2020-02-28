package view

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"main.main/src/utils"
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
	filename = filepath.Clean(utils.CHADIR + "/" + filename)
	filename, err := filepath.Abs(filename)
	if err != nil {
		context.String(http.StatusBadRequest, "invaild path")
		return
	}
	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	if !utils.PathChecker(filename) {
		context.String(http.StatusForbidden, "invaild path")
		return
	}

	fi, err := os.Stat(filename)
	if err != nil {
		context.String(http.StatusNotFound, "The file does not exist")
		return
	}

	if fi.IsDir() {
		files, err := ioutil.ReadDir(filename)
		Check(err)

		dirs := []string{}
		chas := []string{}

		for _, f := range files {
			fi, err := os.Stat(filename + "/" + f.Name())
			Check(err)

			if fi.IsDir() {
				dirs = append(dirs, f.Name())
			} else if filepath.Ext(f.Name()) == ".cha" {
				chas = append(chas, f.Name())
			}
		}

		description := ""
		_, err = os.Stat(filename + "/description.json")
		if err == nil {
			content, err := ioutil.ReadFile(filename + "/description.json")
			if err != nil {
				context.String(http.StatusInternalServerError, "error when reading description")
				return
			}
			description = string(content)
		}

		context.JSON(http.StatusOK, gin.H{"folders": dirs, "files": chas, "description": description})
		return
	} else {
		var dat, err = ioutil.ReadFile(filename)
		Check(err)
		if err != nil {
			context.String(http.StatusInternalServerError, "error when reading file")
			return
		}
		context.JSON(http.StatusOK, gin.H{"content": string(dat)})
		return
	}
}
