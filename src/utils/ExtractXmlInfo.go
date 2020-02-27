package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func extractCell(bs []byte) []interface{} {
	regCell, err := regexp.Compile(`<Cell><Data ss:Type="(.*)">(.*)</Data></Cell>`)

	if err != nil {
		println(err.Error())
	}
	strs := regCell.FindAll(bs, -1)
	println("found:", len(strs))

	ret := make([]interface{}, len(strs))

	for i, str := range strs {
		s := string(str)
		s = strings.Replace(s, `<Cell><Data ss:Type="`, "", 1)
		s = strings.Replace(s, "</Data></Cell>", "", 1)
		s = strings.Replace(s, `">`, " ", 1)
		d := strings.Split(s, " ")
		if d[0] == "String" {
			ret[i] = d[1]
		} else if d[0] == "Number" {
			ret[i], err = strconv.ParseFloat(d[1], 64)
			if err != nil {
				println("error  in conversion:", err.Error())
			}
		}
	}
	return ret
}

//ExtractXMLInfo extract xml contents and process it
func ExtractXMLInfo(bs []byte) [][]interface{} {
	regRow, err := regexp.Compile(`(?s)<Row>(.*?)</Row>`)
	if err != nil {
		println("error in init row extraction")
	}
	//res := [][]interface{}
	rows := regRow.FindAll(bs, -1)
	res := make([][]interface{}, len(rows))
	fmt.Printf("found %d row\n", len(rows))
	for i, r := range rows {
		println("line", i)
		res[i] = extractCell(r)
	}
	return res
}
