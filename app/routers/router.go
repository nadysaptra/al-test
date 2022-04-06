package routers

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App) {
	app.Use(logger.New())

	// Main godoc
	// @Summary Show main route
	// @Description get main route.
	// @Tags root
	// @Accept */*
	// @Produce json
	// @Success 200 "ok"
	// @Router / [get]
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("ok")
	})

	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		return c.SendFile("docs/swagger.json")
	})

	app.Get("/api-docs/*", swagger.New(swagger.Config{ // custom
		URL:          "http://localhost:3000/swagger.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	User(app)
	Event(app)
	Payment(app)
	Ticket(app)
}
