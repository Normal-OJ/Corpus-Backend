package utils

import "regexp"

//CheckOpt check if options are right
func CheckOpt(cmd string, opts string) bool {
	match, _ := regexp.MatchString("(mlu \\+(t\\*\\S+|z\\d+u|f)|freq (\\+t\\*\\S+|\\+s\\S+|\\+u|-t)|mlt \\+t\\*\\S+|kwal (\\+t\\*\\S+|\\+w\\d+|-w\\d+)|kideval)$", cmd+" "+opts)
	return match
}
