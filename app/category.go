package app

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Slug       string `gorm:"not null;unique"`
	Title      string `gorm:"not null"`
	PageType   PageType
	PageTypeID int `gorm:"not null"`
}

type Categories []Category
