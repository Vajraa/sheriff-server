package router

import (
	"sheriff-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/v1/auth")
	auth.Get("/login/github/", handlers.GithubLogin)
	auth.Get("/login/github/callback", handlers.GithubLoginCallback)
}