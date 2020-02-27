package utils

import (
	"os"
	"path/filepath"
)

var CHADIR = os.Getenv("CHA_DIR")
var CHACACHE = os.Getenv("CHA_CACHE")

//PathChecker is an util for checking vaild path
func PathChecker(path string) bool {
	p, err := filepath.Rel(CHADIR, path)
	if err != nil {
		return false
	}

	if filepath.IsAbs(p) {
		return false
	}
	return true
}

//ChaCachePathChecker is a checker for checking cache request path
func ChaCachePathChecker(path string) bool {
	p, err := filepath.Rel(CHACACHE, path)
	if err != nil {
		return false
	}

	if filepath.IsAbs(p) {
		return false
	}
	return true
}
