package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"spe_test/config"
	"spe_test/controller"
	"spe_test/repository"
	"spe_test/service"
)

func main() {
	configuration := config.New()

	// setup repository
	transactionrepo := repository.NewTransactionRepository()
	merchantRepo := repository.NewMerchantpository()
	customerRepo := repository.NewCustomerRepository()

	// setup service
	transactionService := service.NewTransactionService(&transactionrepo, configuration)
	merchantService := service.NewMerchantService(&merchantRepo)
	customerService := service.NewCustomerService(&customerRepo)

	//setup controller
	transactionControllerr := controller.NewTransactionController(&transactionService)
	merchantController := controller.NewMerchantController(&merchantService)
	customerController := controller.NewCustomerController(&customerService)

	//auth login
	authController := controller.NewLoginOauth()

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	//app.Use(config.NewDB())

	app.Use(logger.New())

	// setup routing
	transactionControllerr.Route(app)
	merchantController.Route(app)
	customerController.Route(app)
	authController.Route(app)

	err := app.Listen(":8080")
	if err != nil {
		return
	}

}
