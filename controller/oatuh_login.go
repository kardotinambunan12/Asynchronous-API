package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"spe_test/middleware"
)

type LoginOauth struct {
}

func NewLoginOauth() LoginOauth {
	return LoginOauth{}
}

func (controller *LoginOauth) Route(app *fiber.App) {

	app.Get("/login", controller.LoginHandler)

	app.Get("/auth/google/callback", controller.CallbackUrl)

}

func (controller *LoginOauth) LoginHandler(c *fiber.Ctx) error {
	url := middleware.GoogleOauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func (controller *LoginOauth) CallbackUrl(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Kode OAuth tidak ditemukan, pastikan sudah login melalui Google",
		})
	}

	token, err := middleware.GoogleOauth2Config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error saat Exchange token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mendapatkan token: " + err.Error(),
		})
	}

	client := middleware.GoogleOauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info")
	}
	defer resp.Body.Close()

	var userData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to decode user data")
	}

	return c.JSON(fiber.Map{
		"access_token":  token.AccessToken,
		"token_type":    token.TokenType,
		"expiry":        token.Expiry,
		"refresh_token": token.RefreshToken,
		"userData":      userData,
	})
}
