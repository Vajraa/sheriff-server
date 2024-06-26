package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func GetGithubAccessToken(code string) string {

	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
    
    clientID := GetGithubClientID()
    clientSecret := GetGithubClientSecret()

	requestBodyMap := map[string]string{
        "client_id":     clientID,
        "client_secret": clientSecret,
        "code":          code,
    }
    requestJSON, _ := json.Marshal(requestBodyMap)
    req, reqerr := http.NewRequest(
        "POST",
        "https://github.com/login/oauth/access_token",
        bytes.NewBuffer(requestJSON),
    )
    if reqerr != nil {
        log.Panic("Request creation failed")
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    resp, resperr := http.DefaultClient.Do(req)
    if resperr != nil {
        log.Panic("Request failed")
    }
    defer resp.Body.Close()

    respbody, _ := io.ReadAll(resp.Body)

    var ghresp githubAccessTokenResponse
    json.Unmarshal(respbody, &ghresp)

    return ghresp.AccessToken
}

func GetGithubData(access_token string) string {
    req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
    if err != nil {
        slog.Error("API Request creation failed")
    }

    authorizationHeader := fmt.Sprintf("token %s", access_token)
    req.Header.Set("Authorization", authorizationHeader)

    response, resperr := http.DefaultClient.Do(req)

    if resperr != nil {
        slog.Error("Request Failed")
    }
    defer response.Body.Close()

    respBody, _ := io.ReadAll(response.Body)

    return string(respBody)
}

func GetGithubClientID() string {

	githubClientID := os.Getenv("CLIENT_ID")
	if githubClientID == "" {
		slog.Error("Github Client ID not defined in .env file")
	}
	return githubClientID
}

func GetGithubClientSecret() string {
	
	githubClientSecret := os.Getenv("CLIENT_SECRET")
	if githubClientSecret == "" {
		slog.Error("Github Client Secret not defined in .env file")
	}
	return githubClientSecret
}