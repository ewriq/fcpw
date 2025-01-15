package Routes

import (
	Handler "fcpw/Handler"

	"github.com/gofiber/fiber/v2"
)

func Files(app fiber.Router) {
	app.Post("/list", Handler.Home)
	app.Post("/add", Handler.Home)
	app.Post("/del", Handler.Home)
}