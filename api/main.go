package main

import (
	Handler "fcpw/Handler"
	Middleware "fcpw/Middleware"
	Routes "fcpw/Routes"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(Middleware.Cors)
	app.Use(Middleware.RateLimit)
	app.Use(helmet.New())
	app.Use(Middleware.Compress)

	file := app.Group("/api/file/")

	app.Get("/", Handler.Home)

	Routes.Files(file)

	app.Use(Middleware.Notfound)

	app.Listen(":3000")
}
