package app

import "github.com/jinzhu/gorm"

type Page struct {
	gorm.Model
	PageType   PageType   `gorm:"ForeignKey:PageTypeID;not null"`
	PageTypeID int        `json:"page_type_id"`
	Categories []Category `gorm:"many2many:page_categories;"`
	Tags       []Tag      `gorm:"many2many:page_tags;"`
	Title      string     `gorm:"not null"`
	Slug       string     `gorm:"not null;unique"`
	Content    string     `gorm:"type:text;not null"`
}

type Pages []Page

func (p *Page) IsValid() bool {
	if p.PageTypeID > 0 && len(p.Title) > 0 && len(p.Slug) > 0 {
		return true
	} else {
		return false
	}
}
