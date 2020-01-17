package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
func mltRequestHandler(context *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusBadRequest, "request body missing")
			return
		}
	}()

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
	var cmd *exec.Cmd
	if multi == true {
		//  extract then execute(WIP)
		Unzip(folderName+filename, folderName)
		cmd = exec.Command("./unix-clan/unix/bin/mlt", opts, "*.cha")
	} else {
		cmd = exec.Command("unix-clan/unix/bin/mlt", opts, folderName+"/"+filename)
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
func main() {
	engine := gin.Default()
	engine.POST("/api/mlt", mltRequestHandler)
	engine.Run()
}
