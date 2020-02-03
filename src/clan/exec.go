package clan

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"

	"github.com/gin-gonic/gin"
)

func Exec(c *gin.Context, prog, progOpts, chaSrc string) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *os.PathError:
				c.String(http.StatusNotFound, "Cammand not found.")
			default:
				c.String(http.StatusInternalServerError, "Some other error occurred.")
			}
		}
		return
	}()

	pathToClan := os.Getenv("CLANG_CMD_FOLDER") + "/"

	cmd := exec.Command(pathToClan + prog, progOpts, chaSrc)
	cmd.Stdin = os.Stdin
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	if stdout.Len() != 0 {
		c.String(http.StatusOK, stdout.String())
	} else {
		re := regexp.MustCompile(`Output file <(.+)>`)
		sub := re.FindSubmatch(stderr.Bytes())
		if sub != nil {
			xls := string(sub[1])
			c.FileAttachment(path.Dir(chaSrc) + "/" + xls, xls)
		} else {
			c.String(http.StatusBadRequest, stderr.String())
		}
	}

	return
}
