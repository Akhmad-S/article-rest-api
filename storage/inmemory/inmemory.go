package inmemory

import "github.com/uacademy/article/models"

type InMemory struct {
	Db *Database
}

type Database struct {
	InMemoryArticleData []models.Article
	InMemoryAuthorData  []models.Author
}
