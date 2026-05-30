package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/rizal-mujahiddan/inventoryManagementSystem/backend/pkg/config"
	"github.com/rizal-mujahiddan/inventoryManagementSystem/backend/pkg/database"
)

func main() {

	config.LoadEnv()

	err := database.Connect(
		config.Get("DATABASE_URL"),
	)

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
