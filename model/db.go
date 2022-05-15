package model

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

// データベース接続
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("env error")
	}

	pass:=os.Getenv("PASSWORD")

	dsn := "root:"+pass+"!@tcp(127.0.0.1:5432)/study_go_member?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database can't connect")
	}
	DB.AutoMigrate(&User{})
}