package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Content string
	Category string
	Tags []string
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
		panic("Failed to create a new element in the database !")
	}
	return newPost
}