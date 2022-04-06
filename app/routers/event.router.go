package routers

import (
	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/controllers"
)

func Event(app *fiber.App) {
	event := app.Group("/events")

	event.Get("/", controllers.FetchAllEvents)
	event.Get("/:id", controllers.FetchEvent)
	event.Get("/:id/tickets", controllers.FetchEventTickets)
	event.Post("/", controllers.CreateEvent)
	event.Patch("/:id", controllers.UpdateEvent)
	event.Delete("/:id", controllers.DeleteEvent)
}
