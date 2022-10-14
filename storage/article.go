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
	article.Content = input.Content

	author, err := ReadAuthorById(input.AuthorId)
	if err != nil {
		return err
	}
	article.AuthorId = author.Id
	article.Created_at = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

func ReadArticleById(id string) (models.PackedArticleModel, error) {
	var res models.PackedArticleModel
	for _, v := range InMemoryArticleData {
		if v.Id == id && v.Deleted_at != nil {
			return models.PackedArticleModel{}, errors.New("article already deleted")
		}
		if v.Id == id && v.Deleted_at == nil {
			author, err := ReadAuthorById(v.AuthorId)
			if err != nil {
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
	for _, v := range InMemoryArticleData {
		if v.Deleted_at == nil {
			list = append(list, v)
		}
	}
	return list, err
}

func UpdateArticle(input models.UpdateArticleModel) error {
	var article models.Article
	for i, v := range InMemoryArticleData {
		if v.Id == input.Id && v.Deleted_at == nil {
			article = v
			t := time.Now()
			article.Updated_at = &t
			article.Content = input.Content
			article.AuthorId = input.AuthorId
			InMemoryArticleData[i] = article
			return nil
		}
	}
	return errors.New("article not found")
}

func DeleteArticle(id string) error {
	for i, v := range InMemoryArticleData {
		if v.Id == id {
			t := time.Now()
			v.Deleted_at = &t
			InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}
