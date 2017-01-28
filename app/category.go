package app

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Slug       string
	Title      string
	PageType   PageType
	PageTypeID int
}

type Categories []Category
