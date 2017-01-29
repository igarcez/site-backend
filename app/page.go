package app

import "github.com/jinzhu/gorm"

type Page struct {
	gorm.Model
	PageType   PageType `gorm:"ForeignKey:PageTypeID;not null"`
	PageTypeID int
	Categories []Category `gorm:"many2many:page_categories;"`
	Tags       []Tag      `gorm:"many2many:page_tags;"`
	Title      string     `gorm:"not null"`
	Slug       string     `gorm:"not null;unique"`
	Content    string     `gorm:"type:text;not null"`
}

type Pages []Page
