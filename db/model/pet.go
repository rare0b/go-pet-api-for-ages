package model

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name       string
	Status     string
	Categories []Category `gorm:"many2many:user_categories;"`
	PhotoUrls  []PhotoUrl
	Tags       []Tag `gorm:"many2many:user_tags;"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

type PhotoUrl struct {
	gorm.Model
	PetId uint
	Url   string
}

type Tag struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}
