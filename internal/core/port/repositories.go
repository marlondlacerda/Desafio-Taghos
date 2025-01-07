package port

import "desafio_taghos/internal/core/domain"

type BookRepository interface {
	Create(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book) (*domain.Book, error)
	GetByID(id string) (*domain.Book, error)
	GetAll() ([]*domain.Book, error)
	Delete(id string) error
}
