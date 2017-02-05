package app

import "github.com/igarcez/site-backend/data"

func Init() {
	InitSchemas()
}

func InitSchemas() {
	db := data.GetConnection()
	db.AutoMigrate(&PageType{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Page{})
	db.AutoMigrate(&Tag{})
}
