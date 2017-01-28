package app

import "github.com/jinzhu/gorm"

type Page struct {
	gorm.Model
	PageType   PageType `gorm:"ForeignKey:PageTypeID"`
	PageTypeID int
	Categories []Category `gorm:"many2many:page_categories;"`
	Tags       []Tag      `gorm:"many2many:page_tags;"`
	Title      string
	Slug       string
	Content    string `gorm:"type:text"`
}

type Pages []Page
