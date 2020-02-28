package utils

import (
	"os"
	"path/filepath"
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
	m, err := filepath.Match(CHADIR+"/*", path)
	if err != nil && path != CHADIR {
		return false
	}
	return m || path == CHADIR
}

//ChaCachePathChecker is a checker for checking cache request path
func ChaCachePathChecker(path string) bool {
	path, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	println("checking path:", path)
	m, err := filepath.Match(CHACACHE+"/*", path)
	if err != nil && path != CHACACHE {
		return false
	}
	return m || path == CHACACHE
}
