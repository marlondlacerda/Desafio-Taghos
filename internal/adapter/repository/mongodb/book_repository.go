package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

func (r *BookMongoRepository) GetCollection() *mongo.Collection {
	return r.collection
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": book.ID}
	update := bson.M{"$set": book}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After) // Retorna o documento após a atualização

	var updatedBook domain.Book
	err := r.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedBook)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("document not found")
		}
		return nil, err
	}

	return &updatedBook, nil
}

func (r *BookMongoRepository) GetByID(id string) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var book domain.Book
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("document not found")
		}
		return nil, err
	}

	return &book, nil
}

func (r *BookMongoRepository) GetAll() ([]*domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []*domain.Book
	for cursor.Next(ctx) {
		var book domain.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookMongoRepository) Delete(id string) error {
	return nil
}
