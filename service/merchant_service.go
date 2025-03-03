package service

import (
	"github.com/gofiber/fiber/v2"
	"spe_test/model"
	"spe_test/model/request"
)

type MerchantService interface {
	CreateMerchant(ctx *fiber.Ctx, r *request.Merchant) (*model.WebResponse, error)
	ListMerchant(ctx *fiber.Ctx) (*model.WebResponse, error)
	GetDataMerchantId(ctx *fiber.Ctx, merchantId string) (*model.WebResponse, error)
}
