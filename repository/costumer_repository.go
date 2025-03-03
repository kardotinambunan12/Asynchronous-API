package repository

import (
	"github.com/gofiber/fiber/v2"
	"spe_test/model"
	"spe_test/model/request"
)

type CustomerRepository interface {
	InsertCustomer(ctx *fiber.Ctx, request *request.CustomerRequest, ch chan<- error)
	GetDataCustomer(ctx *fiber.Ctx, ch chan<- *model.WebResponse, errCh chan<- error)
	GetDataCustomerbyId(ctx *fiber.Ctx, request *request.CustomerRequest, ch chan<- *model.WebResponse, errCh chan<- error)
}
