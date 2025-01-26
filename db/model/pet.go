package model

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name       string
	Status     string
	CategoryID uint
	PhotoURLs  []PhotoURL
	Tags       []Tag `gorm:"many2many:user_tags;"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Pets []Pet
}

type PhotoURL struct {
	gorm.Model
	PetID uint
	URL   string
}

type Tag struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}
