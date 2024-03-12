package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/JayeshKarandeDev/fiberFramework/configs"
	"github.com/JayeshKarandeDev/fiberFramework/routes"
	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("Hellow Fibre Service")
	configs.ConfigsLoad()
	app := fiber.New()
	routes.SetRoutes(app)
	// command line arguments
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	log.Println("Server started on http://localhost:", *addr)
	log.Fatal(app.Listen(":" + *addr))
}
