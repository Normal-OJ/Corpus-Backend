package utils

import (
	"hash/crc64"
)

var hf = crc64.New(crc64.MakeTable(crc64.ISO))

//CreateID create a fix length identifier for input string
func CreateID(input string) uint64 {
	hf.Reset()
	hf.Write([]byte(input))
	return hf.Sum64()
}

//CreateFileID does not implemented yet
func CreateFileID(src string) uint64 {
	panic("not implemented :P")
}
