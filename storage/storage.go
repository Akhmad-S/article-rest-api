package storage

import (
	"errors"
	"time"

	"github.com/uacademy/article/models"
)

var InMemoryArticleData []models.Article

func AddArticle(id string, input models.CreateArticleModel) error {
	var article models.Article

	article.Id = id
	article.Author = input.Author
	article.Content = input.Content
	article.Created_at = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

func ReadArticleById(id string) (models.Article, error) {
	for _, v := range InMemoryArticleData {
		if v.Id == id {
			return v, nil
		}
	}
	return models.Article{}, errors.New("article not found")
}

func ReadListArticle() (list []models.Article, err error) {
	list = InMemoryArticleData
	return list, err
}

func UpdateArticle(input models.UpdateArticleModel) (models.Article, error) {
	var article models.Article
	for i, v := range InMemoryArticleData {
		if v.Id == input.Id {
			article = v
			t := time.Now()
			article.Updated_at = &t
			article.Content = input.Content
			article.Author = input.Author
			InMemoryArticleData[i] = article
			return InMemoryArticleData[i], nil
		}
	}
	return models.Article{}, errors.New("article not found")
}

func DeleteArticle(id string) (models.Article, error) {
	for i, v := range InMemoryArticleData {
		if v.Id == id {
			InMemoryArticleData = append(InMemoryArticleData[:i], InMemoryArticleData[i+1:]...)
			return v, nil
		}
	}
	return models.Article{}, errors.New("article not found")
}
