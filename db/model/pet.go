package model

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name       string
	Status     string
	Categories []Category
	PhotoUrls  []PhotoUrl
	Tags       []Tag
}

type Category struct {
	gorm.Model
	PetId uint
	Name  string `gorm:"uniqueIndex"`
}

type PhotoUrl struct {
	gorm.Model
	PetId uint
	Url   string
}

type Tag struct {
	gorm.Model
	PetId uint
	Name  string `gorm:"uniqueIndex"`
}
