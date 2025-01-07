package handler

import (
	"github.com/gofiber/fiber/v2"

	"desafio_taghos/internal/adapter/handler/http/responses"
	"desafio_taghos/internal/core/domain"
	"desafio_taghos/internal/core/service"
	"desafio_taghos/internal/framework/util"

	_ "desafio_taghos/docs"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// @Summary Criar um novo livro
// @Description Endpoint para cadastrar um novo livro no sistema, incluindo título, categoria, autor e sinopse.
// @Tags Book
// @Accept json
// @Produce json
// @Param request body domain.BookCreatedRequest true "query params"
// @Success 201 {object} responses.CreatedResponse
// @Failure 400 {object} responses.BadRequestResponse
// @Failure 500 {object} responses.InternalServerErrorResponse
// @Router /book [post]
func (h *BookHandler) CreateBook(ctx *fiber.Ctx) error {
	book := domain.Book{}

	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "Falha ao processar os dados fornecidos. Verifique o formato da requisição e tente novamente.",
		})
	}

	createdBook, err := h.service.CreateBook(&book)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InternalServerErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(responses.CreatedResponse{
		Message: createdBook.ID,
		Success: true,
		Data:    createdBook,
	})
}

// @Summary Atualizar um livro
// @Description Endpoint para atualizar os dados de um livro existente no sistema.
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "ID do livro"
// @Param request body domain.Book true "Dados atualizados do livro"
// @Success 200 {object} responses.SuccessBookResponse
// @Failure 400 {object} responses.BadRequestResponse
// @Failure 404 {object} responses.NotFoundResponse
// @Failure 500 {object} responses.InternalServerErrorResponse
// @Router /book/{id} [post]
func (h *BookHandler) UpdateBook(ctx *fiber.Ctx) error {
	// Obtendo o ID do livro da URL
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "ID do livro não fornecido.",
		})
	}

	// Validando se o ID é um UUID válido
	if !util.IsValidUUID(id) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "ID do livro fornecido não é um UUID válido.",
		})
	}

	// Parsing dos dados do corpo da requisição
	book := domain.Book{}
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "Falha ao processar os dados fornecidos. Verifique o formato da requisição e tente novamente.",
		})
	}

	// Atualizando o livro
	updatedBook, err := h.service.UpdateBook(id, &book)
	if err != nil {
		if err.Error() == "document not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(responses.NotFoundResponse{
				Success: false,
				Message: "Livro não encontrado.",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InternalServerErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	// Retornando sucesso com os dados atualizados
	return ctx.Status(fiber.StatusOK).JSON(responses.SuccessBookResponse{
		Success: true,
		Message: "Livro atualizado com sucesso.",
		Data:    updatedBook,
	})
}

// @Summary Buscar um livro pelo ID
// @Description Endpoint para buscar um livro existente no sistema pelo ID.
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "ID do livro"
// @Success 200 {object} responses.SuccessBookResponse
// @Failure 400 {object} responses.BadRequestResponse
// @Failure 404 {object} responses.NotFoundResponse
// @Failure 500 {object} responses.InternalServerErrorResponse
// @Router /book/{id} [get]
func (h *BookHandler) GetBookByID(ctx *fiber.Ctx) error {
	// Obtendo o ID do livro da URL
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "ID do livro não fornecido.",
		})
	}

	// Validando se o ID é um UUID válido
	if !util.IsValidUUID(id) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "ID do livro fornecido não é um UUID válido.",
		})
	}

	// Buscando o livro pelo ID
	book, err := h.service.GetBookByID(id)
	if err != nil {
		if err.Error() == "document not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(responses.NotFoundResponse{
				Success: false,
				Message: "Livro não encontrado.",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InternalServerErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	// Retornando sucesso com os dados do livro
	return ctx.Status(fiber.StatusOK).JSON(responses.SuccessBookResponse{
		Success: true,
		Message: "Livro encontrado com sucesso.",
		Data:    book,
	})
}

// @Summary Buscar todos os livros
// @Description Endpoint para buscar todos os livros existentes no sistema.
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {object} responses.SuccessBooksResponse
// @Failure 500 {object} responses.InternalServerErrorResponse
// @Router /book [get]
func (h *BookHandler) GetAllBooks(ctx *fiber.Ctx) error {
	// Buscando todos os livros
	books, err := h.service.GetAllBooks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InternalServerErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	// Retornando sucesso com os dados dos livros
	return ctx.Status(fiber.StatusOK).JSON(responses.SuccessBooksResponse{
		Success: true,
		Message: "Livros encontrados com sucesso.",
		Data:    books,
	})
}

// @Summary Deletar um livro
// @Description Endpoint para deletar um livro existente no sistema pelo ID.
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "ID do livro"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.BadRequestResponse
// @Failure 500 {object} responses.InternalServerErrorResponse
// @Router /book/{id} [delete]
func (h *BookHandler) DeleteBook(ctx *fiber.Ctx) error {
	// Obtendo o ID do livro da URL
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "ID do livro não fornecido.",
		})
	}

	// Validando se o ID é um UUID válido
	if !util.IsValidUUID(id) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequestResponse{
			Success: false,
			Message: "ID do livro fornecido não é um UUID válido.",
		})
	}

	// Deletando o livro
	err := h.service.DeleteBook(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InternalServerErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	// Retornando sucesso
	return ctx.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
		Success: true,
		Message: "Livro deletado com sucesso.",
	})
}
