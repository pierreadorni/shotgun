package main

import (
	"awesomeProject/models"
	"awesomeProject/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// Migrate the schema
	db, err := models.GetDatabase()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	err = models.Migrate(db)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	app.Use(recover.New())
	app.Use(logger.New())

	routes.Register(app)

	app.Listen(":3000")
}
