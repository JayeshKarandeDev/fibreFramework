package main

import (
	"fmt"
	"log"

	"github.com/JayeshKarandeDev/fiberFramework/routes"
	"github.com/gofiber/fiber"
)

func main() {
	fmt.Print("Hellow")
	app := fiber.New()
	routes.SetRoutes(app)
	log.Println("Server started on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
