package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"time"

	"sheriff-server/database"
	"sheriff-server/models"
	"sheriff-server/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GithubLogin(c *fiber.Ctx) error {
	githubClientID := utils.GetGithubClientID()
	redirectURL := fmt.Sprintf(
	"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",githubClientID, "http://localhost:3000/v1/auth/login/github/callback")

	c.Redirect(redirectURL)
	return nil
}


type GithubData struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterUsername   string `json:"twitter_username"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func GithubLoginCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	githubAccessToken := utils.GetGithubAccessToken(code)

	githubData := utils.GetGithubData(githubAccessToken)
	
	var parsedData GithubData
	if err := json.Unmarshal([]byte(githubData), &parsedData); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	user := models.User{
		ID:        primitive.NewObjectID(),
		Username:  parsedData.Login,
		AvatarURL: parsedData.AvatarURL,
		Date:      time.Now().Format(time.RFC3339),
	}

	coll := database.GetCollection("users")
	_, err := coll.InsertOne(c.Context(), user)
	if err != nil {
		slog.Error("Error inserting document:", err)
		return c.Status(500).JSON(fiber.Map{"error": "cannot insert document"})
	}

	log.Println("Inserted document:", user)
	return c.Status(200).JSON(user)
}