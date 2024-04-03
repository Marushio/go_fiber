package main

import (
	"go_fiber/database"
	"go_fiber/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	v1 := app.Group("/api/v1")
	v1.Get("/users", handlers.GetUsers)
	v1.Post("/createUser", handlers.CreateUser)

	log.Fatal(app.Listen(":3000"))

}
