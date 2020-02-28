package modify

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"main.main/src/db"
	"main.main/src/utils"
)

// File is like a file
type File struct {
	Name    string
	Content string
}

// Upload uploads a file
func Upload(file multipart.File, filename string) error {
	dirName := filepath.Dir(filename)
	os.MkdirAll(dirName, os.ModePerm)
	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		out.Close()
	}()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	info, err := utils.ExtractChaID(filename)
	if err != nil {
		return err
	}

	id, err := utils.CreateFileID(filename)

	err = db.InsertFile(id, info.Speaker, info.Age, info.Gender)

	if err != nil {
		return err
	}

	tags := utils.ExtractTag(filename)
	var newTags = make([]string, 0)
	tagIDMap, err := db.QueryTagID()
	if err != nil {
		print(err.Error())
	}

	for _, tag := range tags {
		_, ok := tagIDMap[tag]
		if ok {
			newTags = append(newTags, tag)
		}
	}

	err = db.InsertTag(newTags)

	if err != nil {
		print(err.Error())
	}

	var tagIDs = make([]int64, len(tags))
	for index, tag := range tags {
		tagIDs[index] = utils.CreateID(tag)
	}

	db.InsertRelation(id, tagIDs)
	return nil
}

// UploadRequestHandler is like what it said :P
func UploadRequestHandler(context *gin.Context) {
	filename := context.Query("file")
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}

	filename = filepath.Clean(utils.CHADIR + "/" + filename)

	if !utils.PathChecker(filename) {
		context.JSON(http.StatusBadRequest, gin.H{"result": "unallowed path"})
		return
	}

	err = Upload(file, filename)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	context.JSON(http.StatusOK, gin.H{"result": "success"})
}

// EditRequestHandler is like what it said :P
func EditRequestHandler(context *gin.Context) {
	filename := context.Query("file")
	var file File

	defer func() {
		err := recover()
		if err != nil {
			context.String(http.StatusInternalServerError, "internal server error")
			return
		}
	}()

	context.ShouldBind(&file)

	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": "can't edit file"})
		return
	}

	_, err = os.Stat(file.Name)
	if filename != file.Name && err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": "can't rename file"})
		return
	}
	f.WriteString(file.Content)
	os.Rename(filename, file.Name)

	context.JSON(http.StatusOK, gin.H{"result": "success"})
}

// DeleteRequestHandler is like what it said :P
func DeleteRequestHandler(context *gin.Context) {
	filename := context.Query("file")
	err := os.Remove(filename)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": "can't delete file"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"result": "success"})
}
