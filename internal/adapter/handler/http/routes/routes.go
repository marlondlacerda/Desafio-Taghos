package routes

import (
	"github.com/gofiber/fiber/v2"

	"desafio_taghos/infra/database"
	handler "desafio_taghos/internal/adapter/handler/http"
	repository "desafio_taghos/internal/adapter/repository/mongodb"
	"desafio_taghos/internal/core/service"
)

// BookRouter gerencia as rotas
// @Summary Rotas
// @Description Rotas para gerenciamento dos Books
func BookRouter(app fiber.Router) {
	db := database.MongoDB

	// Injeção de dependências
	bookRepo := repository.NewBookMongoRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	app.Post("/book/create", bookHandler.CreateBook)
}
