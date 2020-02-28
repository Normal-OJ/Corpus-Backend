package utils

import (
	"os"
	"path/filepath"
	"strings"
)

var CHADIR = os.Getenv("CHA_DIR")
var CHACACHE = os.Getenv("CHA_CACHE")

//PathChecker is an util for checking vaild path
func PathChecker(path string) bool {
	path, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	println("checking path:", path)
	i := strings.Index(path, CHADIR)
	println("index:", i)
	return i == 0
}

//ChaCachePathChecker is a checker for checking cache request path
func ChaCachePathChecker(path string) bool {
	path, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	println("checking cache path:", path)
	i := strings.Index(path, CHACACHE)
	println("index:", i)
	return i == 0
}
