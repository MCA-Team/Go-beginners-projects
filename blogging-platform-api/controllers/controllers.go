package controllers

import (
	"blogging-platform-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostOneArticle(c *gin.Context) {
		var inputData models.Article
		c.Bind(&inputData)
		article := models.CreatePost(models.DB, inputData.Title, inputData.Content, inputData.Category, inputData.Tags)

		c.JSON(http.StatusOK, gin.H{
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

