package utils

import (
	"hash/crc64"
	"io/ioutil"
)

var hf = crc64.New(crc64.MakeTable(crc64.ISO))

//CreateID create a fix length identifier for input string
func CreateID(input string) int64 {
	hf.Reset()
	hf.Write([]byte(input))
	ret := int64(hf.Sum64())
	return ret
}

//CreateFileID create an unique identifier base on its content
func CreateFileID(src string) (int64, error) {
	bs, err := ioutil.ReadFile(src)
	if err != nil {
		println("error in creating file id open file:", err.Error())
		return 0, err
	}
	return CreateID(string(bs)), nil
}
