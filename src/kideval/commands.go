package kideval

import (
	"io/ioutil"
	"os"
	"strings"

	"main.main/src/utils"
)

var libraryPath string = os.Getenv("MOR_LIB")
var cmdFolder string = os.Getenv("CLANG_CMD_FOLDER")

/*
	mor execute mor command and overwrite the outFile as result

	[NOTE]: mor command DOES NOT SUPPORT either relative or absolute path(inFile param)
*/
func mor(inFile string, outFile string) {
	libopt := "-l" + libraryPath
	taropt := inFile

	res := utils.RunCmd(cmdFolder+"/mor", []string{libopt, taropt})
	res = res[strings.Index(res, "@UTF8"):]
	println("raw res:")
	println(res)
	ioutil.WriteFile(outFile, []byte(res), 0666)
}
