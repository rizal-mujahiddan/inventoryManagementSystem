package main

import (
	"os"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/rizal-mujahiddan/inventoryManagementSystem/backend/pkg/config"
	"github.com/rizal-mujahiddan/inventoryManagementSystem/backend/pkg/database"
	"fmt"
)

func main() {

	config.LoadEnv()

	dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    config.Get("DB_HOST"),
    config.Get("DB_USER"),
    config.Get("DB_PASSWORD"),
    config.Get("DB_NAME"),
    config.Get("DB_PORT"),
	)

	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_PORT:", os.Getenv("DB_PORT"))
	log.Println("DB_USER:", os.Getenv("DB_USER"))
	log.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	log.Println("DB_NAME:", os.Getenv("DB_NAME"))
	log.Printf("DSN: %s\n", dsn)
	err := database.Connect(dsn)

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
