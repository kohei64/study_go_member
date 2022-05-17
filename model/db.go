package model

import (
    "fmt"
    "os"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB
var err error

// データベース接続
func init() {

    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_DATABASE"),
    )

    // 稼働待ち
    for i := 0; i < 10; i++ {
        DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
						DB.AutoMigrate(&User{})
            fmt.Printf(os.Getenv("DATABASE_URL") + "### connect.\n")
            break
        }
        fmt.Printf("### failed to connect database. connect again.\n")
        time.Sleep(3 * time.Second)
    }

}