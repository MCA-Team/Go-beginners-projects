package controllers

import (
	"blogging-platform-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.String(200, "pong")
  })
  return router
}

func postA(router *gin.Engine) *gin.Engine {
  router.POST("/posts", func(c *gin.Context) {
    var article models.Article
    c.BindJSON(&article)
    c.JSON(http.StatusCreated, article)
  })
  return router
}


func TestPostOneArticle(t *testing.T) {

	t.Run("Test Ping", func(t *testing.T) {
		router := setupRouter()
		
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})

	t.Run("Test of POST method for '/posts' endpoint", func(t *testing.T) {
		router := setupRouter()
		// router.POST("/posts", PostOneArticle)
		// fmt.Println("RTTR")
		
		w := httptest.NewRecorder()
		// Create an example user for testing
		exampleArticle := models.Article{
			Title: "Harry Potter",
			Content:   "Harry Potter was a witch...The end",
			Category: "Thriller",
			Tags: []string{"tag1", "tag2"},
		}
		
		ArticleJson, _ := json.Marshal(exampleArticle)
		req, _ := http.NewRequest("POST", "/posts", strings.NewReader(string(ArticleJson)))
		fmt.Println(req.Host)
		
		fmt.Println(w.Code)
		router.ServeHTTP(w, req)
		fmt.Println(w.Body.String())
		assert.Equal(t, 201, w.Code)
		// Compare the response body with the json data of exampleUser
		assert.Equal(t, string(ArticleJson), w.Body.String())
  })


}

