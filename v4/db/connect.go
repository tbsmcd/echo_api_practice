package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Connect db につなぐ
func Connect() *gorm.DB {
	user := "user"
	password := "pass"
	dbname := "echo_practice"

	connect := user + ":" + password + "@tcp(127.0.0.1:3333)/" + dbname + "?parseTime=true"
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	return db
}
