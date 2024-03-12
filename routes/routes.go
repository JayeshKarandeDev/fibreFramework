package routes

import "github.com/gofiber/fiber"

func home(c *fiber.Ctx) {
	c.Fasthttp.WriteString("Welcome to Simple Fibre ")
}

func SetRoutes(app *fiber.App) {
	app.Get("/", home)
	empRoutes(app)
}
