package zhoseg

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"main.main/src/utils"
)

var zhosegFilePath string = os.Getenv("ZHOSEG_FILE_PATH")
var segment = "seg"
var mode = "tt"

func Call(inputFilePath string) {
	inputFileName := filepath.Base(inputFilePath) //extract only the input file name from the absolute direction
	outputDir, err := os.MkdirTemp("", "ex-zhoseg")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.RemoveAll(outputDir)
	tempFilePath, err := os.MkdirTemp("", "temp-zhoseg")
	if err != nil {
		log.Fatal(err)
		return
	}
	tempFolderPath := filepath.Join(tempFilePath, inputFileName) //generate a temparary folder direction
	fin, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fin.Close()
	fout, err := os.Create(tempFolderPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fout.Close()
	_, err = io.Copy(fout, fin) //copy inputFile to the temp directiry
	if err != nil {
		log.Fatal(err)
		return
	}
	opt := []string{"zhoseg.js", mode, segment, tempFilePath, outputDir}
	//str := strings.Join(command, " ")

	parent := filepath.Dir(inputFilePath)                               //get the parent filepath
	parent = filepath.Base(parent)                                      //extract the parent filename
	outputPath := filepath.Join(outputDir, tempFilePath, inputFileName) //add parent filename to filepath output filepath
	utils.RunCmdDir("node", opt, zhosegFilePath)
	//log.Fatal(opt)

	///modify the segment result
	//////////////////////////////
	data, err := os.OpenFile(outputPath, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	key, _err := os.OpenFile(zhosegFilePath+"/no.txt", os.O_RDWR, 0755)
	if _err != nil {
		log.Fatal(err)
	}

	key2, _err := os.OpenFile(zhosegFilePath+"/ours.txt", os.O_RDWR, 0755)
	if _err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(zhosegFilePath + "/modified.cha")
	if err != nil {
		log.Fatal(err)
	}

	question := []string{}
	word := []string{}

	//question
	scanner_ := bufio.NewScanner(key)
	scanner_.Split(bufio.ScanLines)
	var text_1 []string

	for scanner_.Scan() {
		text_1 = append(text_1, scanner_.Text())
	}

	for _, each_ln_1 := range text_1 {
		question = append(question, each_ln_1)
		//fmt.Println(question)
	}

	//word
	_scanner := bufio.NewScanner(key2)
	_scanner.Split(bufio.ScanLines)
	var _text []string

	for _scanner.Scan() {
		_text = append(_text, _scanner.Text())
	}

	for _, _each_ln2 := range _text {
		word = append(word, _each_ln2)
	}

	//data
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	buffer := ""
	buffer2 := ""
	for _, each_ln := range text {
		buffer = "" //reset buffer
		buffer2 = ""
		//question
		for row, _ := range question {
			if strings.Contains(each_ln, question[row]) && strings.Contains(each_ln, "?") {
				//if strings contains 是不是、對不對、好不好、要不要的問句
				key := []rune(question[row]) //cut the keyword
				str := []rune(each_ln)       //cut the sentence
				buffer = ""                  //clear the buffer
				for i := 0; i < len(str); i++ {
					if str[i] == key[0] && str[i+1] == key[1] && str[i+2] == key[2] && i <= len(str)-2 {
						buffer += string(key[0]) //buffer use to store the new segmented sentence
						buffer += string(" ")
						buffer += string(key[1])
						buffer += string(" ")
						buffer += string(key[2])
						buffer += string(" ")
						i += 2
					} else {
						buffer += string(str[i])
					}
				}
				each_ln = buffer //update the sentence if the sentence contains at least two keywords
			}
		}
		//word
		for row2, _ := range word {
			if strings.Contains(each_ln, word[row2]) {
				//if strings contains 掉、到、們
				str2 := []rune(each_ln) //cut the sentence
				buffer2 += string(str2[0])
				for i := 1; i <= len(str2)-1; i++ {
					if str2[i-1] != ' ' && string(str2[i]) == word[row2] {
						buffer2 += string(" ")
						buffer2 += string(str2[i])
					} else {
						buffer2 += string(str2[i])
					}
				}
				each_ln = buffer2
			}

		}
		f.WriteString(each_ln)
		f.WriteString("\n")
	}

	//os.Rename("modified.cha", pwd+"/programming/test/27_noseg.cha")

	utils.MoveFile(zhosegFilePath+"/modified.cha", outputPath)
	data.Close()
	key.Close()
	key2.Close()
	///////////////////////////////
	e := utils.MoveFile(outputPath, inputFilePath)
	if e != nil {
		log.Fatal(e)
	}

}

func UploadSegmentHandler(context *gin.Context) {
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "file not found"})
		return
	}

	dirName, err := os.MkdirTemp("", "temp")
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer os.RemoveAll(dirName)

	filename := filepath.Join(dirName, "inputData.cha")
	out, err := os.Create(filename)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return
	}
	Call(filename)

	//context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base((filename))))
	//context.Writer.Header().Add("Content-Type", "plain/text; charset=\"utf-8\"")
	context.File(filename)
	//context.String(http.StatusOK, "OK")
}
