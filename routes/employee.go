package routes

import (
	"github.com/JayeshKarandeDev/fiberFramework/controllers"
	"github.com/gofiber/fiber"
)

func empRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
}
