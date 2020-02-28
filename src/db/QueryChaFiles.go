package db

import (
	"database/sql"
	"fmt"
	"strings"
)

var chaQueryCmd = `select distinct cha.path from cha
inner join map on cha.cha_id == map.cha_id
inner join context on context.context_id == map.context_id
%s;`
var ageFilter = "cha.age between ? and ?"
var sexFilter = "cha.sex == ?"
var contextFilter = "context.name in (%s)"

func constructQueryCmd(agesCount int, sexsCount int, contextsCount int) string {
	var opts = []string{}

	// construct age filter
	if agesCount > 0 {
		opts = append(
			opts,
			"("+ageFilter+strings.Repeat(
				fmt.Sprintf(" OR %s", ageFilter), agesCount-1,
			)+")",
		)
	}

	// construct sex filter
	if sexsCount > 0 {
		opts = append(
			opts,
			"("+sexFilter+strings.Repeat(
				fmt.Sprintf(" OR %s", sexFilter), sexsCount-1,
			)+")",
		)
	}
	// construct context filter
	if contextsCount != 0 {
		var o = "?" + strings.Repeat(",?", contextsCount-1)
		opts = append(opts, "("+fmt.Sprintf(contextFilter, o)+")")
	}
	// construct whole cmd
	var condition = ""
	if len(opts) != 0 {
		condition = "where %s"
		c := "%s" + strings.Repeat(" AND %s", len(opts)-1)
		i := make([]interface{}, len(opts))
		for ct := 0; ct != len(opts); ct++ {
			i[ct] = opts[ct]
		}
		condition = fmt.Sprintf(fmt.Sprintf(condition, c), i...)
	}
	var fullCmd = fmt.Sprintf(chaQueryCmd, condition)
	return fullCmd
}

//QueryChaFiles querys db to find the target file base on the condition it given
func QueryChaFiles(ages [][]int, sexs []int, contexts []string) []string {

	database, err := GetDBIns()
	if err != nil {
		println("QueryCha init db error:" + err.Error())
		return nil
	}

	i := make([]interface{}, len(ages)*2+len(sexs)+len(contexts))

	for ct := 0; ct != len(ages); ct++ {
		i[ct*2] = ages[ct][0]
		i[ct*2+1] = ages[ct][1]
	}
	for ct := 0; ct != len(sexs); ct++ {
		i[len(ages)*2+ct] = sexs[ct]
	}
	for ct := 0; ct != len(contexts); ct++ {
		i[len(ages)*2+len(sexs)+ct] = contexts[ct]
	}
	qStr := constructQueryCmd(len(ages), len(sexs), len(contexts))
	println("QueryString:", qStr)
	rows, err := database.Query(qStr, i...)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		println("Exception happend When QueryChaFiles queries DB")
		println("err:", err.Error())
	}
	result := []string{}
	defer rows.Close()
	for rows.Next() {
		var fileName string
		err = rows.Scan(&fileName)
		if err != nil {
			// handle this error
			panic(err)
		}
		result = append(result, fileName)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return result
}
