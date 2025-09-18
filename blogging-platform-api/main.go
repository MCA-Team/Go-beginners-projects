package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"blogging-platform-api/models"
	// "net/http"
	// "github.com/gin-gonic/gin"
)


 

func main() {
	fmt.Println("BLOGGING PLATFORM API")

	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the dabase")
	}
	db.AutoMigrate(&models.Post{})
	

	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// 	})
	// })

	// router.GET("/ap", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 	"apellidos": "Perez Casado",
	// 	})
	// })
	// router.Run(":3000") // listen and serve on 0.0.0.0:3000

	
}