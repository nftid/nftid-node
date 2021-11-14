package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nftid/nftid-node/router"
)

func main() {
	app := fiber.New()

	/*	envs, err := godotenv.Read()

		if err != nil {
			log.Fatal("Error loading .env file")
		}*/

	router.SetupRoutes(app)

	app.Listen(":7535")
}
