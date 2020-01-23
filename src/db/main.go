package db

// TODO: import the issue of muti-threading
import (
	"crypto"
	"database/sql"
	"fmt"
	"time"
)

//UserInfo is the type that stores user info
type UserInfo struct {
	// the type of user
	// admin for "admin"
	// user for "user"
	userType string

	//name of  user
	userName string

	//password of  user
	passwd string

	//the date that user  was created
	date string
}

//UserNotFoundErr defined the error when query failed
type UserNotFoundErr struct {
	target string
}

func (e *UserNotFoundErr) Error() string {
	return fmt.Sprintf("can not found user %s", e.target)
}

var dbRef *sql.DB = nil
var hashFunc = crypto.SHA3_512.New()

var initTableCmd = `
CREATE TABLE IF NOT EXISTS userinfo(
	type TEXT NULL,
	username TEXT NULL,
	passwd TEXT NULL,
	created TEXT NULL
);
`
var addUserCmd = "INSERT INTO userinfo(type,username,passwd,created) VALUES (?,?,?,?);"
var modifyUserPasswdCmd = "UPDATE userinfo SET passwd=? WHERE username=?"
var deleteUserCmd = `yee`
var findUserCmd = "SELECT * FROM userinfo WHERE username=? AND passwd=?"

// hash str and get hex formated string as output
func hash(str string) string {
	hashFunc.Reset()
	hashFunc.Write([]byte(str))
	return fmt.Sprintf("%x", hashFunc.Sum(nil))
}

//InitDB initialize database to the specify form.
func InitDB(database *sql.DB) bool {
	_, err := database.Exec(initTableCmd)
	if err != nil {
		print("error in db func CreateTable:", err.Error())
		return false
	}
	dbRef = database
	return true
}

var checkUserNameExistedCmd = "SELECT * FROM userinfo WHERE username=?"

func checkUserNameExisted(userName string) bool {
	stm, err := dbRef.Prepare(checkUserNameExistedCmd)
	if err != nil {
		print("error in db checkUserNameExisted preparing statement:", err.Error())
	}
	_, err = stm.Exec(userName)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		print("error in db checkUserNameExisted executing query:", err.Error())
		return false
	}
	return true
}

//AddUser add  one user by hashing passwd and executing the sql command(non-checking).
func AddUser(userType string, userName string, passwd string) bool {
	if checkUserNameExisted(userName) {
		return false
	}
	data := UserInfo{
		userType: userType,
		userName: userName,
		passwd:   passwd,
		date:     time.Now().Format("20060102150405"),
	}
	data.passwd = hash(data.passwd)
	stm, err := dbRef.Prepare(addUserCmd)
	if err != nil {
		print("error in db func AddUser:", err.Error())
		return false
	}
	stm.Exec(data)
	return true
}

//FindUser checked whether user exsisted.
func FindUser(username string, passwd string) (*UserInfo, error) {
	passwd = hash(passwd)
	rows, err := dbRef.Query(findUserCmd, username, passwd)
	if err == sql.ErrNoRows {
		return &UserInfo{}, &UserNotFoundErr{fmt.Sprintf("{username:%s,passwd:%s}", username, passwd)}
	} else if err != nil {
		print("error in db FindUser:", err.Error())
		return &UserInfo{}, err
	}
	r := UserInfo{}
	rows.Scan(&r.userType, &r.userName, &r.passwd, &r.date)
	return &r, nil
}

//ModifyUserPasswd modify user info by passing the vaule of modification
func ModifyUserPasswd(username string, passwd string, newpasswd string) bool {
	userinfo, err := FindUser(username, passwd)
	if err != nil {
		print("error in db func ModifyUserPasswd finding matched user:", err.Error())
		return false
	}
	userinfo.passwd = hash(newpasswd)
	stm, err := dbRef.Prepare(modifyUserPasswdCmd)
	if err != nil {
		print("error in db func ModifyUserPasswd preparing statement:", err.Error())
		return false
	}
	_, err = stm.Exec(userinfo.passwd, username)
	if err != nil {
		print("error in db func ModifyUserPasswd modifying password:", err.Error())
		return false
	}
	return true
}

//DeleteUser delete specify user
func DeleteUser(name string, passwd string) bool {
	panic("not implemented :P")
}
