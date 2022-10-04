package main

import(
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	
)

func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}

//CreateArticle
func CreateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New()
	article.Id = id.String()

	article.Created_at = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	c.JSON(http.StatusOK, gin.H{
		"message": "create article",
		"data":    article,
	})
}

//GetArticleById
func GetArticleById(c *gin.Context) {
	id := c.Param("id")

	for _, v := range InMemoryArticleData {
		if v.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "get article by id",
				"data":    v,
			})

			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "article not found",
		"data":    nil,
	})
}

//GetArticleList
func GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get article list",
		"data":    InMemoryArticleData,
	})
}

//UpdateArticle
func UpdateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, v := range InMemoryArticleData {
		if v.Id == article.Id {

			article.Created_at = v.Created_at

			t := time.Now()
			article.Updated_at = &t

			InMemoryArticleData[i] = article

			c.JSON(http.StatusOK, gin.H{
				"message": "update article",
				"data":    article,
			})

			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "article not found",
		"data":    nil,
	})
}

//DeleteArticle
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	for i, v := range InMemoryArticleData {
		if v.Id == id {
			InMemoryArticleData = remove(InMemoryArticleData, i)
			c.JSON(http.StatusOK, gin.H{
				"message": "delete article",
				"data":    v,
			})

			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "article not found",
		"data":    nil,
	})
}
