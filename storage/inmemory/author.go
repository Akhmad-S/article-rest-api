package inmemory

import (
	"github.com/uacademy/article/models"

	"errors"
	"strings"
	"time"
)

func (im InMemory) AddAuthor(id string, input models.CreateAuthorModel) error {
	var author models.Author

	author.Id = id
	author.Firstname = input.Firstname
	author.Lastname = input.Lastname
	author.Created_at = time.Now()

	im.Db.InMemoryAuthorData = append(im.Db.InMemoryAuthorData, author)

	return nil
}

func (im InMemory) ReadAuthorById(id string) (models.Author, error) {
	var res models.Author
	for _, v := range im.Db.InMemoryAuthorData {
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

func (im InMemory) ReadListAuthor(offset, limit int, search string) (list []models.Author, err error) {
	off := 0
	count := 0
	for _, v := range im.Db.InMemoryAuthorData {
		if v.Deleted_at == nil && (strings.Contains(v.Firstname, search) || strings.Contains(v.Lastname, search)) {
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

func (im InMemory) UpdateAuthor(input models.UpdateAuthorModel) error {
	var author models.Author
	for i, v := range im.Db.InMemoryAuthorData {
		if v.Id == input.Id && v.Deleted_at == nil {
			author = v
			t := time.Now()
			author.Updated_at = &t
			author.Firstname = input.Firstname
			author.Lastname = input.Lastname
			im.Db.InMemoryAuthorData[i] = author
			return nil
		}
	}
	return errors.New("author not found")
}

func (im InMemory) DeleteAuthor(id string) error {
	for i, v := range im.Db.InMemoryAuthorData {
		if v.Id == id {
			t := time.Now()
			v.Deleted_at = &t
			im.Db.InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("author not found")
}
