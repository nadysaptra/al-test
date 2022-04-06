package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App) {
	app.Use(logger.New())

	// ShowAccount godoc
	// @Summary      Show an account
	// @Description  get string by ID
	// @Tags         accounts
	// @Accept       json
	// @Produce      json
	// @Param        id   path      int  true  "Account ID"
	// @Success      200  {object}  model.Account
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError
	// @Router       /accounts/{id} [get]
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("ok")
	})

	User(app)
	Event(app)
	Payment(app)
	Ticket(app)
}
