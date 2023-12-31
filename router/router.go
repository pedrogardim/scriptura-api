package router

import (
	"scriptura/scriptura-api/handler"
	m "scriptura/scriptura-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app fiber.Router) {
	api := app.Group("/api/v1", logger.New())

	api.Use(m.JsonMiddleware)

	bible := api.Group("/bible", m.PaginationMiddleware)

	bible.Get("/book/:book", handler.GetBible)
	bible.Get("/book/:book/chapter/:chapter", handler.GetBible)
	bible.Get("/book/:book/chapter/:chapter/verse/:start-:end", handler.GetBible)
	bible.Get("/book/:book/chapter/:chapter/verse/:start", handler.GetBible)
}
