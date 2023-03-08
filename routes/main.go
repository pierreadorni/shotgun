package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Register(app *fiber.App) {
	store = session.New(session.Config{
		Expiration: 24 * 60 * 60 * 1000, // 24 hours
	})

	// front and metrics
	app.Static("/", "./front/dist")
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Awesome Project Metrics"}))

	// login and logout
	app.Get("/login", validateTicket, loginCas)
	app.Get("/logout", logout)
	app.Get("/session", GetSession)

	// api routes
	api := app.Group("/api")
	api.Get("/events", getEvents)
	api.Get("/events/:id", getEvent)
	api.Post("/events", createEvent)
	api.Put("/events/:id", updateEvent)
	api.Delete("/events/:id", deleteEvent)
	api.Post("/events/:id/subscribe", subscribe)
}
