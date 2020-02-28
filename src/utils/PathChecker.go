package utils

import (
	"os"
	"path/filepath"
)

var CHADIR = os.Getenv("CHA_DIR")
var CHACACHE = os.Getenv("CHA_CACHE")

//PathChecker is an util for checking vaild path
func PathChecker(path string) bool {
	println("checking path:", path)
	p, err := filepath.Rel(CHADIR, path)
	if err != nil {
		return false
	}
	println("rel p:", p)
	println("is abs:", filepath.IsAbs(p))
	return filepath.IsAbs(p)
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
