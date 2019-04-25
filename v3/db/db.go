package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
}

// Init は db を初期化するメソッド
func Init() {
	user := "user"
	password := "pass"
	dbname := "echo_practice"

	connect := user + ":" + password + "@tcp(127.0.0.1:3333)/" + dbname
	db, err = gorm.Open("mysql", connect)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

}
