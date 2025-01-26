package model

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name       string
	Status     string
	Categories []Category `gorm:"many2many:user_categories;"`
	PhotoURLs  []PhotoURL
	Tags       []Tag `gorm:"many2many:user_tags;"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

type PhotoURL struct {
	gorm.Model
	PetId uint
	URL   string
}

type Tag struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}
