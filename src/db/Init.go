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
var addSample = `
INSERT INTO cha(cha_id, path, age, sex)
VALUES (1, 'a.cha', 12, 1);

INSERT INTO cha(cha_id, path, age, sex)
VALUES (2, 'b/b.cha', 24, 2);

INSERT INTO context(context_id, name)
VALUES (1, 'toy');

INSERT INTO context(context_id, name)
VALUES (2, 'book');

INSERT INTO map(cha_id, context_id)
VALUES (1, 1);

INSERT INTO map(cha_id, context_id)
VALUES (1, 2);

INSERT INTO map(cha_id, context_id)
VALUES (2, 2);
`
var db_ins *sql.DB = nil

func Init(database *sql.DB) error {
	if db_ins != nil {
		return nil
	}
	db_ins = database
	_, err := database.Exec(dbSchema)
	if err != nil {
		println("error when creating schema:", err.Error())
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
func AddTestSample() {
	database, err := GetDBIns()
	if err != nil {
		println("Error in init AddTestSample:", err.Error())
		return
	}
	_, err = database.Exec(addSample)
	if err != nil {
		println("error in AddTestSample adding test sample:", err.Error())
	}
}
