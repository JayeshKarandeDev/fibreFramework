package routes

import (
	"github.com/JayeshKarandeDev/fiberFramework/controllers"
	"github.com/gofiber/fiber"
)

func accRoutes(app *fiber.App) {
	app.Get("/accounts", controllers.GetAccounts)
}
