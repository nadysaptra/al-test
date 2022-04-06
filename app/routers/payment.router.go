package routers

import (
	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/controllers"
)

func Payment(app *fiber.App) {
	payment := app.Group("/payment")

	payment.Get("/", controllers.FetchAllPayments)
	payment.Get("/:id", controllers.FetchPayment)
	payment.Patch("/:id", controllers.UpdatePayment)
}
