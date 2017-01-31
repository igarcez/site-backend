package app

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Slug       string `gorm:"not null;unique"`
	Title      string `gorm:"not null"`
	PageType   PageType
	PageTypeID int `gorm:"not null" json:"page_type_id"`
}

type Categories []Category

func (category *Category) IsValid() bool {
	if len(category.Slug) > 0 &&
		len(category.Title) > 0 &&
		category.PageTypeID > 0 {
		return true
	} else {
		return false
	}
}
