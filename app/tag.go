package app

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Tag string `gorm:"not null;unique"`
}

type Tags []Tag
