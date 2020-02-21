package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var idPattern = "@ID:\t.*?"

//IDinfo is for recording info inside ID header
type IDinfo struct {
	lang      string
	database  string
	Speaker   string
	Age       int
	Gender    int
	group     string
	ses       string
	character string
	edu       string
	other     string
}

//ErrAstolfo used  for undefined gender
var ErrAstolfo = errors.New("undefine  gender , it might be Astolfo")

func getAge(str string) (int, error) {
	str = strings.ReplaceAll(str, ".", "")
	ages := strings.Split(str, ";")
	if len(ages) != 2 {
		return 0, strconv.ErrSyntax
	}
	y, err := strconv.Atoi(ages[0])
	if err != nil {
		return 0, err
	}
	m, err := strconv.Atoi(ages[1])
	if err != nil {
		return 0, err
	}
	return y*12 + m, nil
}
func getGender(str string) (int, error) {
	if str == "male" {
		return 0, nil
	} else if str == "female" {
		return 1, nil
	} else {
		return 2, ErrAstolfo
	}
}

//ExtractChaID is used to extract info inside cha file id header
func ExtractChaID(fileSrc string) (IDinfo, error) {
	file, err := os.Open(fileSrc)
	if err != nil {
		log.Fatal(err)
		return IDinfo{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(idPattern)
	for scanner.Scan() {
		txt := scanner.Text()
		if !reg.Match([]byte(txt)) {
			continue
		}

		txt = strings.Replace(txt, "@ID:\t", "", 1)
		raw := strings.Split(txt, "|")
		if len(raw) != 10 && raw[2] != "CHI" {
			continue
		}
		age, err := getAge(raw[3])
		if err != nil {
			return IDinfo{}, err
		}
		gender, err := getGender(raw[4])
		if err != nil {
			return IDinfo{}, err
		}
		ret := IDinfo{
			lang:      raw[0],
			database:  raw[1],
			Speaker:   raw[2],
			Age:       age,
			Gender:    gender,
			group:     raw[5],
			ses:       raw[6],
			character: raw[7],
			edu:       raw[8],
			other:     raw[9],
		}
		return ret, nil
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return IDinfo{}, err
	}
	return IDinfo{}, nil
}
