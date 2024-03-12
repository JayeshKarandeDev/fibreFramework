package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/JayeshKarandeDev/fiberFramework/routes"
	"github.com/gofiber/fiber"
)

func main() {
	fmt.Print("Hellow")
	app := fiber.New()
	routes.SetRoutes(app)
	// command line arguments
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	log.Println("Server started on http://localhost:", *addr)
	log.Fatal(app.Listen(":" + *addr))
}
