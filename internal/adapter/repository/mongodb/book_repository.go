package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"desafio_taghos/internal/core/domain"
	"desafio_taghos/internal/core/port"
)

type BookMongoRepository struct {
	collection *mongo.Collection
}

func NewBookMongoRepository(db *mongo.Database) port.BookRepository {
	return &BookMongoRepository{
		collection: db.Collection("books"),
	}
}

func (r *BookMongoRepository) Create(book *domain.Book) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookMongoRepository) Update(book *domain.Book) (*domain.Book, error) {
	return nil, nil
}

func (r *BookMongoRepository) GetByID(id string) (*domain.Book, error) {
	return nil, nil
}

func (r *BookMongoRepository) GetAll() ([]*domain.Book, error) {
	return nil, nil
}

func (r *BookMongoRepository) Delete(id string) error {
}
