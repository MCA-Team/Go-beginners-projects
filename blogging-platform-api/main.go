package main

import (
	"blogging-platform-api/models"
	"fmt"
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

	// post := models.CreatePost(db, "title", "the body", "the category", []string{"tag1", "tag2", "tag3"})
	// fmt.Println(post)
	
	query := db.Find(&models.Post{})
	fmt.Println(query.RowsAffected)
	router := gin.Default()
	router.POST("/posts", func (c *gin.Context) {
		
		var inputData models.Post
		c.Bind(&inputData)
		post := models.CreatePost(db, inputData.Title, inputData.Content, inputData.Category, inputData.Tags)

		c.JSON(http.StatusOK, gin.H{
		"post": post,
		})
	})


	router.GET("/posts", func (c *gin.Context) {
		
		var posts []models.Post
		db.Find(&posts)

		c.JSON(http.StatusOK, gin.H{
		"post": posts,
		})
	})


	router.Run(":3000") // listen and serve on 0.0.0.0:3000

	
}