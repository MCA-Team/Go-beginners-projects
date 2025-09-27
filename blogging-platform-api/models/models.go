package models

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string	`form:"title" json:"title" xml:"title"`
	Content string	`form:"content" json:"content" xml:"content"`
	Category string	`form:"category" json:"category" xml:"category"`
	Tags []string `form:"tags" json:"tags" xml:"tags" gorm:"serializer:json"`	// Serialization allows the dabase to store slice of stings correctly
}

var DB *gorm.DB

func ConnectToDB(DB_path string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(DB_path), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the dabase")
	} else {
		fmt.Println("DB CONNECTION ESTABLISHED")
	}
	
}
