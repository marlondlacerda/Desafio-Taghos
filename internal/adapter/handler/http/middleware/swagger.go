package middleware

import (
	"desafio_taghos/docs"
	"desafio_taghos/infra/config"
	"strings"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
)

// Swagger é um middleware que serve a interface Swagger UI para a documentação da API.
func Swagger(app *fiber.App) {
	if config.Env("ENVIRONMENT") != "dev" {
		// Caso não seja dev, não serve a documentação
		return
	}

	swaggerJSON := docs.SwaggerInfo.ReadDoc()

	// Serve o arquivo JSON da documentação da API
	app.Get("/docs/swagger.json", func(c *fiber.Ctx) error {
		// Substitui o valor da variável HOST no JSON da documentação
		modifiedJSON := strings.Replace(swaggerJSON, "${HOST}", config.Host, -1)
		return c.SendString(modifiedJSON)
	})

	// Serve a interface Swagger UI para a documentação da API
	app.Get("/swagger/*", func(ctx *fiber.Ctx) error {
		host := ctx.BaseURL()

		html, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: host + "/docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Desafio Taghos API Documentation",
			},
			DarkMode: true,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		ctx.Set("Content-Type", "text/html; charset=utf-8")
		return ctx.SendString(html)
	})
}
