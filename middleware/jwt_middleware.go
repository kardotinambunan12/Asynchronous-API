package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/oauth2/google"
	"spe_test/config"
	"spe_test/model"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	configuration := config.New()

	// Set secret key from .env file.
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:    []byte(configuration.Get("JWT_SECRET_KEY")),
		ContextKey:    "jwt", // used in private routes
		SigningMethod: "HS512",
		ErrorHandler:  jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Message: err.Error(),
			Code:    "400",
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
		Message: err.Error(),
		Code:    "401",
	})
}
