package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/uacademy/article/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/uacademy/article/handlers"
	"github.com/uacademy/article/models"
	"github.com/uacademy/article/storage"
)

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"

	err := storage.AddAuthor("6d7795cd-e91c-4878-a26a-79cb436ab22a", models.CreateAuthorModel{
		Firstname: "John",
		Lastname:  "Doe",
	})
	if err != nil {
		panic(err)
	}

	err = storage.AddArticle("286002a5-0021-45d3-bdb1-c502fa1ef9a9", models.CreateArticleModel{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	//template GET method
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/article", handlers.CreateArticle)
		v1.GET("/article/:id", handlers.GetArticleById)
		v1.GET("/article", handlers.GetArticleList)
		v1.PUT("/article", handlers.UpdateArticle)
		v1.DELETE("/article/:id", handlers.DeleteArticle)

		v1.POST("/author", handlers.CreateAuthor)
		v1.GET("/author/:id", handlers.GetAuthorById)
		v1.GET("/author", handlers.GetAuthorList)
		v1.PUT("/author", handlers.UpdateAuthor)
		v1.DELETE("/author/:id", handlers.DeleteAuthor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
