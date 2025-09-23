package main

import (
	"blogging-platform-api/models"
	"blogging-platform-api/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)


 

func main() {
	fmt.Println("BLOGGING PLATFORM API")

	models.ConnectToDB()
	models.DB.AutoMigrate(&models.Article{})
	
	// query := models.DB.Find(&models.Post{})
	// fmt.Println(query.RowsAffected)
	router := gin.Default()

	router.POST("/posts", controllers.PostOneArticle)
	router.GET("/posts", controllers.GetAllArticles)
	router.GET("/posts/:articleID", controllers.GetSpecificArticle)
	router.DELETE("/posts/:articleID", controllers.DeleteSpecificArticle)
	router.PUT("/posts/:articleID", controllers.UpdateOneElement)


	router.Run(":3000") // listen and serve on 0.0.0.0:3000

	
}