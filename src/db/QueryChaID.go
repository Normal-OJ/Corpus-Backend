package db

import "database/sql"

var queryChaIDFromCha = `select cha_id,path from cha`
var queryChaIDFromChaStmt *sql.Stmt = nil

//QueryChaID query all the exsisted cha and cha ID inside database
func QueryChaID() (map[string]int64, error) {
	if queryChaIDFromChaStmt == nil {
		database, err := GetDBIns()
		if err != nil {
			println("error  when QueryChaID obtain db instance:", err.Error())
			return nil, err
		}
		queryChaIDFromChaStmt, err = database.Prepare(queryChaIDFromCha)
		if err != nil {
			println("error when QueryChaID preparing  statement:", err.Error())
			return nil, err
		}
	}
	rows, err := queryChaIDFromChaStmt.Query()
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	ret := make(map[string]int64)

	for rows.Next() {
		var path string = ""
		var id int64 = 0
		rows.Scan(&id, &path)
		ret[path] = id

		if rows.Err() != nil {
			return nil, rows.Err()
		}
	}
	return ret, nil
}
