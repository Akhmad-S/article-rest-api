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
	article.AuthorId = input.AuthorId
	article.Content = input.Content
	article.Created_at = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

func ReadArticleById(id string) (models.PackedArticleModel, error) {
	var res models.PackedArticleModel
	for _, v := range InMemoryArticleData {
		if v.Id == id {
			author, err := ReadAuthorById(v.AuthorId)
			if err != nil{
				return res, err
			}
			res.Id = v.Id
			res.Content = v.Content
			res.Author = author
			res.Created_at = v.Created_at
			res.Updated_at = v.Updated_at
			res.Deleted_at = v.Deleted_at
			return res, nil
		}
	}
	return models.PackedArticleModel{}, errors.New("article not found")
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
			article.AuthorId = input.AuthorId
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
