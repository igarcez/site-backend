package app

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Tag string
}

type Tags []Tag
