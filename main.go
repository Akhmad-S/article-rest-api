package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	"github.com/uacademy/article/docs" // docs is generated by Swag CLI, you have to import it.
)

var InMemoryArticleData []Article

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"

	InMemoryArticleData = append(InMemoryArticleData, Article{
		Id: "1",
	})

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
