package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// データベース接続
func init() {
	dsn := "root:Mmky1031!@tcp(127.0.0.1:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database can't connect")
	}
}