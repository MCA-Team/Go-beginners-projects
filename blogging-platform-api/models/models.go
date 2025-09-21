package models

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Content string
	Category string
	Tags []string `gorm:"serializer:json"`	// Serialization allows the dabase to store slice of stings correctly
}

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the dabase")
	} else {
		fmt.Println("DB CONNECTION ESTABLISHED")
	}
	
}

// CreatePost creates a new post with passed title, body, category and tags elements in argument and commits it in the DB in argument too through GORM.
func CreatePost(DB *gorm.DB, title, body, category string, tags []string) Article {
	newArticle := Article {
		Title: title,
		Content: body,
		Category: category,
		Tags: tags,
	}
	if result := DB.Create(&newArticle); result.Error != nil {
		panic(result.Error)
	}
	return newArticle
}