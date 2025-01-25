package model

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name      string
	Status    string
	PhotoUrls []PhotoUrl
}

type PhotoUrl struct {
	gorm.Model
	PetId uint
	Url   string
}

type Tag struct {
	gorm.Model
	PetId uint
	Name  string
}
