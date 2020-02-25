package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//Context struct is use for storing the info about  json type string from file
type Context struct {
	Contexts []string `json:"context"`
}

var tagPattern = "@Comment:\t{.*?}"

//ExtractTag extracts tags inside the file
func ExtractTag(fileSrc string) []string {
	file, err := os.Open(fileSrc)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(tagPattern)
	for scanner.Scan() {
		txt := scanner.Text()
		if !reg.Match([]byte(txt)) {
			continue
		}

		txt = strings.Replace(txt, "@Comment:\t", "", 1)
		var filetag Context
		err := json.Unmarshal([]byte(txt), &filetag)
		if err != nil {
			fmt.Println("JsonStrToTag err: ", err)
		}
		return filetag.Contexts
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
