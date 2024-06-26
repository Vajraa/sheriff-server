package handlers

import (
	"fmt"
	"log/slog"

	"sheriff-server/utils"

	"github.com/gofiber/fiber/v2"
)

func GithubLogin(c *fiber.Ctx) error {
	githubClientID := utils.GetGithubClientID()
	redirectURL := fmt.Sprintf(
	"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",githubClientID, "http://localhost:3000/login/github/callback")
	c.Redirect(redirectURL)
	return nil
}

func GithubLoginCallback(c *fiber.Ctx) error { 
	code := c.Query("code")
	githubAccessToken := utils.GetGithubAccessToken(code)

	githubData := utils.GetGithubData(githubAccessToken)

	slog.Info(githubData)

	fmt.Println(githubAccessToken)
	return c.Status(200).JSON(githubData)
}