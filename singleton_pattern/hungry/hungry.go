package hungry

import (
	"database/sql"
	"fmt"
)

var MainDB *sql.DB

func init() {
	MainDB, _ = sql.Open("mysql", "mysqldsn")
	err := MainDB.Ping()
	if err != nil {
		panic("connect mysql failed")
	}
	fmt.Println("connect mysql successfully")
}
