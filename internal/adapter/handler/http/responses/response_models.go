package responses

import "desafio_taghos/internal/core/domain"

type BadRequestResponse struct {
	Success bool   `json:"success" example:"false"`                               // Indica se a criação foi bem-sucedida
	Message string `json:"message" example:"Bad Request when try to create book"` // Mensagem de erro
}

type NotFoundResponse struct {
	Success bool   `json:"success" example:"false"`          // Indica se a criação foi bem-sucedida
	Message string `json:"message" example:"Book not found"` // Mensagem de erro
}

type InternalServerErrorResponse struct {
	Success bool   `json:"success" example:"false"`                 // Indica se a criação foi bem-sucedida
	Message string `json:"message" example:"Internal Server Error"` // Mensagem de erro
}

type CreatedResponse struct {
	Success bool         `json:"success" example:"true"`                      // Indica se a criação foi bem-sucedida
	Message string       `json:"message" example:"Book created successfully"` // Mensagem de sucesso
	Data    *domain.Book `json:"data"`                                        // Dados do livro criado
}

// SuccessResponse representa uma resposta genérica de sucesso
type SuccessResponse struct {
	Success bool   `json:"success" example:"true"`                 // Indica se a operação foi bem-sucedida
	Message string `json:"message" example:"Operation successful"` // Mensagem de sucesso
}

// SuccessBookResponse representa uma resposta de sucesso, com ou sem dados
type SuccessBookResponse struct {
	Success bool         `json:"success" example:"true"`                 // Indica se a operação foi bem-sucedida
	Message string       `json:"message" example:"Operation successful"` // Mensagem de sucesso
	Data    *domain.Book `json:"data,omitempty"`                         // Dados do livro
}

// SuccessBooksResponse representa uma resposta de sucesso com todos os livros
type SuccessBooksResponse struct {
	Success bool           `json:"success" example:"true"`                            // Indica se a operação foi bem-sucedida
	Message string         `json:"message" example:"Livros encontrados com sucesso."` // Mensagem de sucesso
	Data    []*domain.Book `json:"data"`                                              // Lista de livros
}
