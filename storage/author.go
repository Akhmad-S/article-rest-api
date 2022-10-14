package storage

import (
	"github.com/uacademy/article/models"

	"errors"
	"time"
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
		if v.Id == id && v.Deleted_at != nil {
			return models.Author{}, errors.New("author already deleted")
		}
		if v.Id == id && v.Deleted_at == nil {
			res = v
			return res, nil
		}
	}
	return res, errors.New("author not found")
}

func ReadListAuthor() (list []models.Author, err error) {
	for _, v := range InMemoryAuthorData {
		if v.Deleted_at == nil {
			list = append(list, v)
		}
	}
	return list, err
}

func UpdateAuthor(input models.UpdateAuthorModel) error {
	var author models.Author
	for i, v := range InMemoryAuthorData {
		if v.Id == input.Id && v.Deleted_at == nil {
			author = v
			t := time.Now()
			author.Updated_at = &t
			author.Firstname = input.Firstname
			author.Lastname = input.Lastname
			InMemoryAuthorData[i] = author
			return nil
		}
	}
	return errors.New("author not found")
}

func DeleteAuthor(id string) error {
	for i, v := range InMemoryAuthorData {
		if v.Id == id {
			t := time.Now()
			v.Deleted_at = &t
			InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("author not found")
}
