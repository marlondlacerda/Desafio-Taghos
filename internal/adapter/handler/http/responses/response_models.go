package responses

import "desafio_taghos/internal/core/domain"

type BadRequestResponse struct {
	Success bool   `json:"success" example:"false"`                               // Indica se a criação foi bem-sucedida
	Message string `json:"message" example:"bad request when try to create book"` // Mensagem de erro
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
