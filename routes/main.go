package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Register(app *fiber.App) {
	store = session.New(session.Config{CookiePath: ""})

	// front and metrics
	app.Static("/", "./front/dist")
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Awesome Project Metrics"}))

	// api routes
	api := app.Group("/api")
	api.Get("/accounts", getAccounts)
	api.Get("/accounts/:id", visibleByAccountId, getAccount)
	api.Post("/accounts", createAccount)
	api.Put("/accounts/:id", visibleByAccountId, updateAccount)
	api.Delete("/accounts/:id", visibleByAccountId, deleteAccount)

	api.Post("/login", login)
	api.Post("/logout", logout)

}
