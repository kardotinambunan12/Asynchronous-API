package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"spe_test/config"
	"spe_test/model/request"
	"strings"
)

var (
	configuration = config.New()
)

var GoogleOauth2Config = &oauth2.Config{
	ClientID:     configuration.Get("GOOGLE_CLIENT_ID"),
	ClientSecret: configuration.Get("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  configuration.Get("GOOGLE_REDIRECT_URI"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func ValidateGoogleToken(accessToken string) (*request.GoogleTokenInfo, error) {
	url := "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=" + accessToken
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid access token")
	}

	var tokenInfo request.GoogleTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, err
	}

	expectedAudience := configuration.Get("GOOGLE_CLIENT_ID")
	if tokenInfo.Audience != expectedAudience {
		return nil, errors.New("invalid audience")
	}

	return &tokenInfo, nil
}

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization Header",
		})
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization Header format",
		})
	}

	accessToken := parts[1]

	// Validasi token Google
	_, err := ValidateGoogleToken(accessToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid access token",
		})
	}

	return c.Next()
}
