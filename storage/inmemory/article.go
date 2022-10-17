package inmemory

import (
	"errors"
	"strings"
	"time"

	"github.com/uacademy/article/models"
)

func (im InMemory) AddArticle(id string, input models.CreateArticleModel) error {
	var article models.Article
	article.Id = id
	article.Content = input.Content

	author, err := im.ReadAuthorById(input.AuthorId)
	if err != nil {
		return err
	}
	article.AuthorId = author.Id
	article.Created_at = time.Now()

	im.Db.InMemoryArticleData = append(im.Db.InMemoryArticleData, article)

	return nil
}

func (im InMemory) ReadArticleById(id string) (models.PackedArticleModel, error) {
	var res models.PackedArticleModel
	for _, v := range im.Db.InMemoryArticleData {
		if v.Id == id && v.Deleted_at != nil {
			return models.PackedArticleModel{}, errors.New("article already deleted")
		}
		if v.Id == id && v.Deleted_at == nil {
			author, err := im.ReadAuthorById(v.AuthorId)
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

func (im InMemory) ReadListArticle(offset, limit int, search string) (list []models.Article, err error) {
	off := 0
	count := 0
	for _, v := range im.Db.InMemoryArticleData {
		if v.Deleted_at == nil && (strings.Contains(v.Content.Title, search) || strings.Contains(v.Content.Body, search)) {
			if off >= offset {
				count++
				list = append(list, v)
			}
			if count >= limit {
				break
			}
			off++
		}
	}
	return list, err
}

func (im InMemory) UpdateArticle(input models.UpdateArticleModel) error {
	var article models.Article
	for i, v := range im.Db.InMemoryArticleData {
		if v.Id == input.Id && v.Deleted_at == nil {
			article = v
			t := time.Now()
			article.Updated_at = &t
			article.Content = input.Content
			im.Db.InMemoryArticleData[i] = article
			return nil
		}
	}
	return errors.New("article not found")
}

func (im InMemory) DeleteArticle(id string) error {
	for i, v := range im.Db.InMemoryArticleData {
		if v.Id == id {
			t := time.Now()
			v.Deleted_at = &t
			im.Db.InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found")
}
