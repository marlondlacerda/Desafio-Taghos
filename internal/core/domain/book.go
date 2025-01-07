package domain

import (
	"time"

	_ "desafio_taghos/docs"
)

// BookCreatedRequest model info
// @Description Estrutura de dados usada para criar um novo livro, sem incluir campos de controle como ID, CreatedAt e UpdatedAt.
type BookCreatedRequest struct {
	Title    string `json:"title" bson:"title"`       // Título do livro
	Category string `json:"category" bson:"category"` // Categoria do livro
	Author   string `json:"author" bson:"author"`     // Autor do livro
	Synopsis string `json:"synopsis" bson:"synopsis"` // Sinopse do livro
}

// Book model info
// @Description Informações completas do livro, incluindo ID, título, categoria, autor, sinopse, e timestamps de criação e atualização.
type Book struct {
	ID        string    `json:"_id" bson:"_id"`               // ID do livro (gerado automaticamente pelo banco de dados)
	Title     string    `json:"title" bson:"title"`           // Título do livro
	Category  string    `json:"category" bson:"category"`     // Categoria do livro
	Author    string    `json:"author" bson:"author"`         // Autor do livro
	Synopsis  string    `json:"synopsis" bson:"synopsis"`     // Sinopse do livro
	CreatedAt time.Time `json:"created_at" bson:"created_at"` // Data e hora de criação do livro (gerado automaticamente)
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` // Data e hora da última atualização do livro
}
