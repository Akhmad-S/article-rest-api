package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}

// CreateArticle godoc
// @Summary     Create article
// @Description create new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     CreateArticleModel true "Article body"
// @Success     200     {object} JSONResult{data=Article}
// @Failure     400     {object} JSONError
// @Router      /v1/article [post]
func CreateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	id := uuid.New()
	article.Id = id.String()

	article.Created_at = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	c.JSON(http.StatusOK, JSONResult{
		Message: "create article",
		Data:    article,
	})
}

// GetArticle godoc
// @Summary     Get article
// @Description get article by ID
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} JSONResult{data=Article}
// @Failure     404 {object} JSONError
// @Router      /v1/article/{id} [get]
func GetArticleById(c *gin.Context) {
	id := c.Param("id")

	for _, v := range InMemoryArticleData {
		if v.Id == id {
			c.JSON(http.StatusOK, JSONResult{
				Message: "get article by id",
				Data:    v,
			})

			return
		}
	}
	c.JSON(http.StatusNotFound, JSONError{
		Error: "article not found",
	})
}

// ListArticles godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Success     200 {object} JSONResult{data=[]Article}
// @Router      /v1/article [get]
func GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, JSONResult{
		Message: "get article list",
		Data:    InMemoryArticleData,
	})
}

// UpdateArticle godoc
// @Summary     Update article
// @Description update article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     UpdateArticleModel true "Article body"
// @Success     200     {object} JSONResult{data=Article}
// @Failure     400     {object} JSONError
// @Failure     404     {object} JSONError
// @Router      /v1/article [put]
func UpdateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}
	for i, v := range InMemoryArticleData {
		if v.Id == article.Id {

			article.Created_at = v.Created_at

			t := time.Now()
			article.Updated_at = &t

			InMemoryArticleData[i] = article

			c.JSON(http.StatusOK, JSONResult{
				Message: "update article",
				Data:    article,
			})

			return
		}
	}
	c.JSON(http.StatusNotFound, JSONError{
		Error: "article not found",
	})
}

// DeleteArticle godoc
// @Summary     Delete article
// @Description delete article by ID
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} JSONResult{data=Article}
// @Failure     404 {object} JSONError
// @Router      /v1/article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	for i, v := range InMemoryArticleData {
		if v.Id == id {
			InMemoryArticleData = remove(InMemoryArticleData, i)
			c.JSON(http.StatusOK, JSONResult{
				Message: "delete article",
				Data:    v,
			})

			return
		}
	}
	c.JSON(http.StatusNotFound, JSONError{
		Error: "article not found",
	})
}
