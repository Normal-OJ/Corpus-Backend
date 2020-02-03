package route

import (
	"os"
	"main.main/src/clan"

	"github.com/gin-gonic/gin"
)

func clanRequestHandler(c *gin.Context) {
	pathToChat := os.Getenv("CHAT_FOLDER") + "/"

	prog := c.Param("prog")
	progOpts := c.PostForm("opts")
	chaSrc := pathToChat + c.PostForm("path")

	clan.Exec(c, prog, progOpts, chaSrc)

	return
}
