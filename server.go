package main

import (
	"awesomeProject/models"
	"awesomeProject/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	app := fiber.New()

	// load Env vars
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file: " + err.Error())
	}

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

	// add allow origin and allow credentials header for frontend hostname
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONT_ADDR"),
		AllowCredentials: true,
	}))

	routes.Register(app)

	app.Listen(":3000")
}
