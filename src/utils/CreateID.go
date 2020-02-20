package utils

import (
	"hash/crc64"
	"io/ioutil"
)

var hf = crc64.New(crc64.MakeTable(crc64.ISO))

//CreateID create a fix length identifier for input string
func CreateID(input string) uint64 {
	hf.Reset()
	hf.Write([]byte(input))
	return hf.Sum64()
}

//CreateFileID create an unique identifier base on its content
func CreateFileID(src string) (uint64, error) {
	bs, err := ioutil.ReadFile(src)
	if err != nil {
		println("error in creating file id open file:", err.Error())
		return 0, err
	}
	return CreateID(string(bs)), nil
}
