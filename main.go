package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var InMemoryArticleData []Article

func main() {
	InMemoryArticleData = make([]Article, 0)
	r := gin.Default()
	//template GET method
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/article", CreateArticle)
		v1.GET("/article/:id", GetArticleById)
		v1.GET("/article", GetArticleList)
		v1.PUT("/article", UpdateArticle)
		v1.DELETE("/article/:id", DeleteArticle)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
