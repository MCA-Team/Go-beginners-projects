package controllers

import (
	"blogging-platform-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOneArticle(c *gin.Context) {
		var inputData models.Article
		c.Bind(&inputData)
		article := models.CreatePost(models.DB, inputData.Title, inputData.Content, inputData.Category, inputData.Tags)

		c.JSON(http.StatusCreated, gin.H{
		"Article": article,
		})
	}


func GetAllArticles(c *gin.Context) {
		var articles []models.Article
		models.DB.Find(&articles)

		c.JSON(http.StatusOK, gin.H{
		"Articles": articles,
		})
	}


func GetSpecificArticle(c *gin.Context) {
	var article models.Article
	articleID := c.Param("articleID")
	models.DB.First(&article, articleID)	// Retrieval by primary key. ID being the default primary key
	fmt.Println(article)
	c.JSON(http.StatusOK, gin.H{
		"Articles": article,
		})
}


func DeleteSpecificArticle(c *gin.Context) {
	var article models.Article
	articleID := c.Param("articleID")
	models.DB.Delete(&article, articleID)	// Deletion by primary key. ID being the default primary key
	fmt.Println(article)
	c.String(http.StatusNoContent, "Article %v deleted !", articleID)
}


func UpdateOneElement(c *gin.Context) {
	var inputData models.Article
	var postToUpdate models.Article
	articleID := c.Param("articleID")

	c.Bind(&inputData)

	// Finding the article to update
	models.DB.First(&postToUpdate, articleID)

	// Applying updates on the article to update
	models.DB.Model(&postToUpdate).Updates(models.Article{
		Title : inputData.Title,
		Content : inputData.Content,
		Category : inputData.Category,
		Tags : inputData.Tags,
	})
	c.JSON(http.StatusOK, gin.H{
	"Article": postToUpdate,
	})
}
