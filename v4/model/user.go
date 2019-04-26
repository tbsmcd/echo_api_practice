package model

import (
	"../db"
	"github.com/jinzhu/gorm"
)

// User モデルの構造体
type User struct {
	gorm.Model
	Name string
	// 本当は CreaedAt とかも定義したほうが良い
}

// AddUser ユーザの新規追加
func AddUser(name string) error {
	dbconn := db.Connect()
	defer dbconn.Close()

	u := User{}
	u.Name = name
	return dbconn.Create(&u).Error
}

// FindUser　ユーザの取得
func FindUser(id string) (User, error) {
	dbconn := db.Connect()
	defer dbconn.Close()

	u := User{}
	err := dbconn.First(&u, id).Error
	return u, err
}
