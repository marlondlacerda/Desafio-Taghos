package service

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"

	"desafio_taghos/internal/core/domain"
	"desafio_taghos/internal/core/port"
)

type BookService struct {
	repository port.BookRepository
}

func NewBookService(repository port.BookRepository) *BookService {
	return &BookService{repository: repository}
}

func (s *BookService) CreateBook(book *domain.Book) (*domain.Book, error) {
	newUUID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("failed to generate UUID")
	}

	book.ID = newUUID.String()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	return s.repository.Create(book)
}
