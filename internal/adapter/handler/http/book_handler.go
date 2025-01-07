package handler

import (
	"github.com/gofiber/fiber/v2"

	"desafio_taghos/internal/adapter/handler/http/responses"
	"desafio_taghos/internal/core/domain"
	"desafio_taghos/internal/core/service"

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
// @Router /book/create [post]
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
