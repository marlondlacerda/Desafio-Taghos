package routes

import (
	"github.com/gofiber/fiber/v2"

	_ "desafio_taghos/docs"
	"desafio_taghos/infra/database"
	handler "desafio_taghos/internal/adapter/handler/http"
	repository "desafio_taghos/internal/adapter/repository/mongodb"
	"desafio_taghos/internal/core/service"
)

// IndexRouter gerencia as rotas
// @Summary Rotas
// @Description Rotas para gerenciamento da indexação do banco de dados
func IndexingRouter(app fiber.Router) {
	db := database.MongoDB

	// Injeção de dependências
	bookRepo := repository.NewBookMongoRepository(db)
	indexingService := service.NewIndexingService(bookRepo)
	indexingHandler := handler.NewIndexingHandler(indexingService)

	app.Get("/indexing/book", indexingHandler.CreateIndexesBook)
}
