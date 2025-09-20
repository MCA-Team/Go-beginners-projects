package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Content string
	Category string
	Tags []string `gorm:"serializer:json"`	// Serialization allows the dabase to store slice of stings correctly
}


// CreatePost creates a new post with passed title, body, category and tags elements in argument and commits it in the db in argument too through GORM.
func CreatePost(db *gorm.DB, title, body, category string, tags []string) Post {
	newPost := Post {
		Title: title,
		Content: body,
		Category: category,
		Tags: tags,
	}
	if result := db.Create(&newPost); result.Error != nil {
		panic(result.Error)
	}
	return newPost
}