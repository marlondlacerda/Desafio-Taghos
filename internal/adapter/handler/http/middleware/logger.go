package middleware

import "github.com/gofiber/fiber/v2/middleware/logger"

func Logger() logger.Config {
	return logger.Config{
		Format:     "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 - 15:04:05",
		TimeZone:   "America/Sao_Paulo",
	}
}
