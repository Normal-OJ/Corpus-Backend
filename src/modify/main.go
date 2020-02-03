package modify

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// File is like a file
type File struct {
	Name    string
	Content string
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
