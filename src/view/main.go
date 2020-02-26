package view

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var CHADIR = os.Getenv("CHA_DIR")

// Check error
func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func dirChecker(path string) bool {
	p, err := filepath.Rel(CHADIR, path)
	if err != nil {
		return false
	}

	if filepath.IsAbs(p) {
		return false
	}
	return true
}

// RequestHandler is like what it said :P
func RequestHandler(context *gin.Context) {
	filename := context.Query("file")
	filename = filepath.Clean(CHADIR + "/" + filename)
	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	if !dirChecker(filename) {
		context.String(http.StatusForbidden, "invaild path")
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
			}
			description = string(content)
		}

		context.JSON(http.StatusOK, gin.H{"folders": dirs, "files": chas, "description": description})
	} else {
		var dat, err = ioutil.ReadFile(filename)
		Check(err)
		if err != nil {
			context.String(http.StatusInternalServerError, "error when reading file")
		}
		context.JSON(http.StatusOK, gin.H{"content": string(dat)})
	}
}
