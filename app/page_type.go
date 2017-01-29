package app

import "github.com/jinzhu/gorm"

type PageType struct {
	Code        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	gorm.Model
}

type PageTypes []PageType
