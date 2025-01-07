package handler

import (
	"github.com/gofiber/fiber/v2"

	_ "desafio_taghos/docs"
	"desafio_taghos/internal/adapter/handler/http/responses"
	"desafio_taghos/internal/core/service"
)

type IndexingHandler struct {
	service *service.IndexingService
}

func NewIndexingHandler(service *service.IndexingService) *IndexingHandler {
	return &IndexingHandler{service: service}
}

// @Summary Criar índices para a coleção de livros
// @Description Endpoint para criar índices nos campos "author", "category", "synopsis" e "title" da coleção de livros no MongoDB.
// @Tags Indexing
// @Accept json
// @Produce json
// @Success 200 {object} responses.SuccessResponse "Índices criados com sucesso"
// @Failure 500 {object} responses.InternalServerErrorResponse "Erro interno no servidor"
// @Router /indexing/book [get]
func (h *IndexingHandler) CreateIndexesBook(c *fiber.Ctx) error {
	err := h.service.CreateIndexesBook()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
		Success: true,
		Message: "Indexes criado com sucesso",
	})
}
