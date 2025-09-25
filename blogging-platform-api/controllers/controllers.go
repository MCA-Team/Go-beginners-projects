package controllers

import (
	"blogging-platform-api/models"
	"fmt"
	"encoding/json"
	"net/http"
	"errors"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func PostOneArticle(c *gin.Context) {
	var inputData models.Article

	// Strict JSON decoding
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // ❗ Forbid unknown fields

	if err := decoder.Decode(&inputData); err != nil {	// Handling error for unknows fields and incorrect Model binding
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid JSON: %v", err),
		})
		return
	}

	newArticle := models.Article {
		Title: inputData.Title,
		Content: inputData.Content,
		Category: inputData.Category,
		Tags: inputData.Tags,
	}
	if result := models.DB.Create(&newArticle); result.Error != nil {
		panic(result.Error)
	}
	c.JSON(http.StatusCreated, gin.H{
		"Article": newArticle,
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
	
	
	if err := models.DB.First(&article, articleID).Error; errors.Is(err, gorm.ErrRecordNotFound) {	// Retrieval of the article by primary key. ID being the default primary key
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Article with ID '%v' not found", articleID)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Articles": article,
		})
}


func DeleteSpecificArticle(c *gin.Context) {
	var article models.Article
	articleID := c.Param("articleID")

	fmt.Println(article)
	if err := models.DB.First(&article, articleID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Article with ID '%v' not found", articleID)})
		return
	}

	models.DB.Unscoped().Delete(&article, articleID)	// PERMANENT deletion by primary key. ID being the default primary key (Omit `Unscoped()` statement to avoid permanent deletion)
	c.String(http.StatusNoContent, "Article %v deleted !", articleID)	// No body response returned for 204 status 'No content' code
}


func UpdateOneElement(c *gin.Context) {
	var inputData models.Article
	var postToUpdate models.Article
	articleID := c.Param("articleID")
	// Strict JSON decoding
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // ❗ Forbid unknown fields


	if err := models.DB.First(&postToUpdate, articleID).Error; errors.Is(err, gorm.ErrRecordNotFound) {	// Finding the article to update
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err := decoder.Decode(&inputData); err != nil {	// Handling error for unknows fields and incorrect Model binding
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid JSON: %v", err),
		})
		return
	}
	

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
