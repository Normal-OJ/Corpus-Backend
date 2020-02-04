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

func QueryCha(ages []interface{} , sexs []int , contexts []string){
	var opts = []string{}
	// construct age filter
	for i:=0;i!=len(ages);i++{
		opts=append(opts , ageFilter)
	}
	// construct sex filter
	for i:=0;i!=len(sexs);i++{
		opts=append(opts,sexFilter)
	}
	// construct context filter
	if len(contexts) !=0{
		var o = "?" + strings.Repeat(",?" , len(contexts))
		opts=append(opts,fmt.Sprintf(contextFilter , o)) 
	}

	// construct whole cmd
	var condition = ""
	if len(opts) != 0{
		condition = "where %s"
		c := "%s" + strings.Repeat(" OR %s" , len(opts))
		i:= make([]interface{} , len(opts))
		for ct , str := range opts{
			i[ct] = str
		}
		condition = fmt.Sprintf(fmt.Sprintf(condition , c) , i...)
	}
	var fullCmd = fmt.Sprintf(chaQueryCmd , condition)

	database , err = GetDBIns()
}