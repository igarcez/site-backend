package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func GetConnection() *gorm.DB {
	return db
}

func InitDataConnection() {
	var err error

	db, err = gorm.Open("mysql", "testuser:TestPasswd9@/test?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
}

func CloseDataConnection() {
	db.Close()
}
