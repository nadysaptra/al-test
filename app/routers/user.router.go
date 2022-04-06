package routers

import (
	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/controllers"
)

func User(app *fiber.App) {
	user := app.Group("/users")

	user.Get("/", controllers.FetchAllUsers)
	user.Get("/:id", controllers.FetchUser)
	user.Get("/:id/tickets", controllers.FetchUserTickets)
	user.Post("/", controllers.CreateUser)
	user.Patch("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)
}
