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
	"github.com/uacademy/article/storage/inmemory"
)

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"

	var stg storage.StorageI
	stg = inmemory.InMemory{
		Db: &inmemory.Database{},
	}

	err := stg.AddAuthor("6d7795cd-e91c-4878-a26a-79cb436ab22a", models.CreateAuthorModel{
		Firstname: "John",
		Lastname:  "Doe",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("286002a5-0021-45d3-bdb1-c502fa1ef9a9", models.CreateArticleModel{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("46fd0c19-1f65-4b7d-b1c6-b4a7f649a7b1", models.CreateArticleModel{
		Content: models.Content{
			Title: "1",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("de010394-73bb-4913-82c0-6f2743dc0aac", models.CreateArticleModel{
		Content: models.Content{
			Title: "2",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("981397bb-246f-43e0-8cc1-28d664799493", models.CreateArticleModel{
		Content: models.Content{
			Title: "3",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("6e2c4560-5004-4af1-b346-49269c4e0b859", models.CreateArticleModel{
		Content: models.Content{
			Title: "4",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("615ca1e1-76f6-440a-b64a-c8292cd8a901", models.CreateArticleModel{
		Content: models.Content{
			Title: "5",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("fd3df66d-472a-4c4e-bbc0-331d9e945b4a", models.CreateArticleModel{
		Content: models.Content{
			Title: "6",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("85b0d3bf-c1c5-4fec-b9a8-8ac15a5ee18b", models.CreateArticleModel{
		Content: models.Content{
			Title: "7",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("f0d86eda-927b-4a5e-914c-bb8c6802d516", models.CreateArticleModel{
		Content: models.Content{
			Title: "8",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("710d2531-ae9c-4fd6-b55f-ad48cdc42301", models.CreateArticleModel{
		Content: models.Content{
			Title: "9",
			Body:  "Something about lorem",
		},
		AuthorId: "6d7795cd-e91c-4878-a26a-79cb436ab22a",
	})
	if err != nil {
		panic(err)
	}

	err = stg.AddArticle("efd51b13-3e25-45ad-9ae8-d5ce99bc4b23", models.CreateArticleModel{
		Content: models.Content{
			Title: "10",
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

	h := handlers.Handler{
		Stg: stg,
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/article", h.CreateArticle)
		v1.GET("/article/:id", h.GetArticleById)
		v1.GET("/article", h.GetArticleList)
		v1.PUT("/article", h.UpdateArticle)
		v1.DELETE("/article/:id", h.DeleteArticle)

		v1.POST("/author", h.CreateAuthor)
		v1.GET("/author/:id", h.GetAuthorById)
		v1.GET("/author", h.GetAuthorList)
		v1.PUT("/author", h.UpdateAuthor)
		v1.DELETE("/author/:id", h.DeleteAuthor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
