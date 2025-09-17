package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)



func main() {
	fmt.Println("BLOGGING PLATFORM API")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080

	
}