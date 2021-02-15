package kideval

import (
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
func mor(inFile string) {
	libopt := "-L" + libraryPath
	utils.RunCmd(cmdFolder+"/mor", []string{libopt, inFile})
}

func post(inFile string) {
	libopt := "-d" + libraryPath + "/post.db"
	utils.RunCmd(cmdFolder+"/post", []string{libopt, inFile})
}

func postmortem(inFile string) {
	libopt := "-L" + libraryPath
	utils.RunCmd(cmdFolder+"/postmortem", []string{libopt, inFile})
}

func megrasp(inFile string) {
	libopt := "-L" + libraryPath
	utils.RunCmd(cmdFolder+"/megrasp", []string{libopt, inFile})
}

// check checks if a file's format is correct
func check(file string) bool {
	res := utils.RunCmd(cmdFolder+"/check", []string{file})
	println("raw res:")
	println(res)
	return !strings.Contains(res, "File")
}
