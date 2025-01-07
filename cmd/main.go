package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"desafio_taghos/infra/config"
	"desafio_taghos/infra/database"
	"desafio_taghos/internal/adapter/handler/http/middleware"
	"desafio_taghos/internal/framework/logs"
)

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

	return app
}
