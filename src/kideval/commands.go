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
	libopt := "-L" + libraryPath
	taropt := inFile

	res := utils.RunCmd(cmdFolder+"/mor", []string{libopt, taropt})
	res = res[strings.Index(res, "@UTF8"):]
	println("raw res:")
	println(res)
	ioutil.WriteFile(outFile, []byte(res), 0666)
}

func post(inFile string, outFile string) {
	libopt := "-d" + libraryPath + "/post.db"
	taropt := inFile

	res := utils.RunCmd(cmdFolder+"/post", []string{libopt, taropt})
	res = res[strings.Index(res, "@UTF8"):]
	println("raw res:")
	println(res)
	ioutil.WriteFile(outFile, []byte(res), 0666)
}

func postmortem(inFile string, outFile string) {
	libopt := "-L" + libraryPath
	taropt := inFile

	res := utils.RunCmd(cmdFolder+"/postmortem", []string{libopt, taropt})
	res = res[strings.Index(res, "@UTF8"):]
	println("raw res:")
	println(res)
	ioutil.WriteFile(outFile, []byte(res), 0666)
}

func megrasp(inFile string, outFile string) {
	libopt := "-L" + libraryPath
	taropt := inFile

	res := utils.RunCmd(cmdFolder+"/megrasp", []string{libopt, taropt})
	res = res[strings.Index(res, "@UTF8"):]
	println("raw res:")
	println(res)
	ioutil.WriteFile(outFile, []byte(res), 0666)
}

// check checks if a file's format is correct
func check(file string) bool {
	res := utils.RunCmd(cmdFolder+"/check", []string{file})
	println("raw res:")
	println(res)
	return !strings.Contains(res, "File")
}
