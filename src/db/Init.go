package db

import (
	"database/sql"
	"errors"
)

var dbSchema = `CREATE TABLE IF NOT EXISTS cha (
	cha_id integer primary key, 
	path varchar(20), 
	age integer,
	sex integer
  );
  
  CREATE TABLE IF NOT EXISTS context (
	context_id integer primary key, 
	name varchar(20)
  );
  
  CREATE TABLE IF NOT EXISTS map (
	cha_id integer,
	context_id integer,
	FOREIGN KEY(cha_id) REFERENCES cha(id),
	FOREIGN KEY(context_id) REFERENCES context(id)
  );`
var db_ins *sql.DB = nil

func Init(database *sql.DB) error {
	if db_ins != nil {
		return nil
	}
	_, err := database.Exec(dbSchema)
	if err != nil {
		print("error when creating schema:", err.Error())
		return err
	}
	return nil
}

var InstanceNotExisted = errors.New("DB instance not created yet")

func GetDBIns() (*sql.DB, error) {
	if db_ins == nil {
		return nil, InstanceNotExisted
	}
	return db_ins, nil
}
