package app

import "github.com/jinzhu/gorm"

type PageType struct {
	Code        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	gorm.Model
}

type PageTypes []PageType

func (pageType *PageType) IsValid() bool {
	if len(pageType.Code) > 0 && len(pageType.Description) > 0 {
		return true
	} else {
		return false
	}
}
