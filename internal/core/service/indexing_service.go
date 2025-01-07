package service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"desafio_taghos/internal/core/port"
)

type IndexingService struct {
	repository port.BookRepository
}

func NewIndexingService(repository port.BookRepository) *IndexingService {
	return &IndexingService{repository: repository}
}

func (s *IndexingService) CreateIndexesBook() error {
	// Criando índices no MongoDB
	indexes := []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"author": 1}, // Índice para o campo 'author'
			Options: options.Index().SetUnique(false),
		},
		{
			Keys:    map[string]interface{}{"category": 1}, // Índice para o campo 'category'
			Options: options.Index().SetUnique(false),
		},
		{
			Keys:    map[string]interface{}{"synopsis": 1}, // Índice para o campo 'synopsis'
			Options: options.Index().SetUnique(false),
		},
		{
			Keys:    map[string]interface{}{"title": 1}, // Índice para o campo 'title'
			Options: options.Index().SetUnique(false),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Criando os índices no banco de dados
	for _, index := range indexes {
		_, err := s.repository.GetCollection().Indexes().CreateOne(ctx, index)
		if err != nil {
			return errors.New("failed to create index: " + err.Error())
		}
	}

	return nil
}
