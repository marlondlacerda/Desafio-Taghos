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

func (s *BookService) UpdateBook(id string, book *domain.Book) (*domain.Book, error) {
	book.UpdatedAt = time.Now()

	updatedBook, err := s.repository.Update(book)
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (s *BookService) GetBookByID(id string) (*domain.Book, error) {
	book, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) GetAllBooks() ([]*domain.Book, error) {
	books, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}
