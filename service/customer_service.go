package service

import (
	"github.com/gofiber/fiber/v2"
	"spe_test/model"
	"spe_test/model/request"
)

type CustomerService interface {
	GetDataCustomer(ctx *fiber.Ctx) (*model.WebResponse, error)
	InsertCustomer(ctx *fiber.Ctx, r *request.CustomerRequest) (*model.WebResponse, error)
	GetDataCustomerbyId(ctx *fiber.Ctx, customerPan string) (*model.WebResponse, error)
}
