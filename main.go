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

	db.AutoMigrate(&Pet{}, &PhotoUrl{}, &Tag{})

	db.Create(&Pet{Name: "Goma Shiba", Status: "available", PhotoUrls: []PhotoUrl{"https://example.com/", "https://example.com/"}, Tags: []Tag{"Inu", "Shiba"}})
}
