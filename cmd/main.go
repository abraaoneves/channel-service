package main

import (
	"log"

	"github.com/abraaoneves/channel/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	handlers.RegisterChannelRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
