package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("petstore.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Pet{}, &Category{}, &PhotoURL{}, &Tag{})

	db.Create(&Pet{Name: "Goma Shiba", Status: "available", Category: Category{"Inu"}, PhotoURLs: []PhotoURL{"https://example.com/", "https://example.com/"}, Tags: []Tag{"Shiba", "Nihonken"}})
}
