package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nftid/nftid-node/router"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":7535")
}
