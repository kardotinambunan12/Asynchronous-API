package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"spe_test/config"
	errorhandler "spe_test/error_handler"
	"strings"
	"time"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64
	Email   string
	Role    string
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	configuration := config.New()
	return []byte(configuration.Get("JWT_SECRET_KEY")), nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			return nil, errors.New("Missing or malformed JWT")
		}
		return nil, err
	}

	return token, nil
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))
		email := string(claims["email"].(string))
		role := string(claims["role"].(string))

		return &TokenMetadata{
			Expires: expires,
			Email:   email,
			Role:    role,
		}, nil
	}

	return nil, err
}

func JwtVerificationWithClaim(ctx *fiber.Ctx) *TokenMetadata {
	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(ctx)

	if err != nil {
		// message := err.Error()
		message := "Maaf, permintaan ditolak token invalid"
		panic(errorhandler.UnauthorizedError{
			message,
		})
	}

	expires := claims.Expires
	if now > expires {
		message := "sesi login telah habis"

		panic(errorhandler.UnauthorizedError{
			Message: message,
		})
	}

	return claims
}
