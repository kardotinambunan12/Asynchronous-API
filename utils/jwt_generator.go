package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"spe_test/config"
	"strconv"
	"time"
)

func GenerateNewAccessToken(email string) (string, error) {
	configuration := config.New()

	// Set secret key from .env file.
	secret := configuration.Get("JWT_SECRET_KEY")
	fmt.Println("secret : ", secret)
	// Set expires hours count for secret key from .env file.
	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRE_HOUR_COUNT"))
	fmt.Println("hoursCount : ", hoursCount)
	// Create a new claims.
	claims := jwt.MapClaims{}
	// Set public claims:
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
