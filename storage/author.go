package storage

import (
	"github.com/uacademy/article/models"
	
	"time"
	"errors"
)

var InMemoryAuthorData []models.Author

func AddAuthor(id string, input models.CreateAuthorModel) error {
	var author models.Author

	author.Id = id
	author.Firstname = input.Firstname
	author.Lastname = input.Lastname
	author.Created_at = time.Now()

	InMemoryAuthorData = append(InMemoryAuthorData, author)

	return nil
}

func ReadAuthorById(id string) (models.Author, error) {
	var res models.Author
	for _, v := range InMemoryAuthorData {
		if v.Id == id {
			res = v
			return res, nil
		}
	}
	return res, errors.New("author not found")
}

func ReadListAuthor() (list []models.Author, err error) {
	list = InMemoryAuthorData
	return list, err
}

func UpdateAuthor(input models.UpdateAuthorModel) (models.Author, error) {
	var author models.Author
	for i, v := range InMemoryAuthorData {
		if v.Id == input.Id {
			author = v
			t := time.Now()
			author.Updated_at = &t
			author.Firstname = v.Firstname
			author.Lastname = v.Lastname
			InMemoryAuthorData[i] = author
			return InMemoryAuthorData[i], nil
		}
	}
	return models.Author{}, errors.New("author not found")
}

func DeleteAuthor(id string) (models.Author, error) {
	for i, v := range InMemoryAuthorData {
		if v.Id == id {
			InMemoryAuthorData = append(InMemoryAuthorData[:i], InMemoryAuthorData[i+1:]...)
			return v, nil
		}
	}
	return models.Author{}, errors.New("author not found")
}