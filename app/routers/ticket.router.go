package routers

import (
	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/controllers"
)

func Ticket(app *fiber.App) {
	ticket := app.Group("/tickets")

	ticket.Get("/", controllers.FetchAllTicket)
	ticket.Post("/", controllers.CreateTicket)
	ticket.Patch("/:id", controllers.UpdateTicket)
}
