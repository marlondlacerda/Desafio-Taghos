package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"desafio_taghos/infra/config"
	"desafio_taghos/infra/database"
	"desafio_taghos/internal/adapter/handler/http/middleware"
	"desafio_taghos/internal/adapter/handler/http/routes"
	"desafio_taghos/internal/framework/logs"

	_ "desafio_taghos/docs"
)

// @title Desafio Taghos
// @version 1.0
// @description API de Book
// @termsOfService http://swagger.io/terms/

// @contact.name API Support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host ${HOST}
// @BasePath /
// @schemes http https
func main() {
	app := Setup()
	log.Fatal(app.Listen(":" + config.Env("PORT")))
}

func Setup() *fiber.App {
	logs.InitializeLogger()
	logs.Logger.Info().Msg("Servidor rodando na porta 3000...")

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  config.ServerName,
	})

	if mongoErr := database.NewMongoDBDatabaseConnection(false); mongoErr != nil {
		log.Fatal(mongoErr)
	}

	middleware.Common(app)
	middleware.Swagger(app)

	routes.BookRouter(app)

	return app
}
