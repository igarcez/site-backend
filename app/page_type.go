package app

import (
	"github.com/igarcez/site-backend/data"
	"github.com/jinzhu/gorm"
)

type PageType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	gorm.Model
}

type PageTypes []PageType

func NewPageType() PageType {
	return PageType{}
}

func (pt *PageType) GetCollection() PageTypes {
	db := data.GetConnection()

	result := PageTypes{}
	db.Find(&result)
	return result
}
