package config

import (
	"github.com/gofiber/fiber/v2"
	errorhandler "spe_test/error_handler"
)

func NewFiberConfig() fiber.Config {

	return fiber.Config{
		ErrorHandler: errorhandler.ErrorHandler,
	}
}
