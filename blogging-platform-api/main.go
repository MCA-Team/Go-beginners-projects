package main

import (
	"blogging-platform-api/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)


 

func main() {
	fmt.Println("BLOGGING PLATFORM API")

	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the dabase")
	}

	db.AutoMigrate(&models.Post{})

	// post := models.CreatePost(db, "title", "the body", "the category", "['tag1', 'tag2', 'tag3']")
	// fmt.Println(post)
	

	router := gin.Default()
	router.POST("/posts", func(c *gin.Context) {
		var inputData struct {
			title string 
			content string 
			category string
			tags string	
		}

		c.Bind(&inputData)

		post := models.CreatePost(db, inputData.title, inputData.content, inputData.category, inputData.tags)
		log.Println(inputData)
		c.JSON(http.StatusOK, gin.H{
		"post": post,
		})
	})

	router.Run(":3000") // listen and serve on 0.0.0.0:3000

	
}