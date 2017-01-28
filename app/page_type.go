package app

import "github.com/jinzhu/gorm"

type PageType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	gorm.Model
}

type PageTypes []PageType
