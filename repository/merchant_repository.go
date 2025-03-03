package repository

import (
	"github.com/gofiber/fiber/v2"
	"spe_test/model"
	"spe_test/model/request"
)

type MerchantRepository interface {
	CreateMerchant(ctx *fiber.Ctx, request *request.Merchant, ch chan<- error)
	ListMerchant(ctx *fiber.Ctx, ch chan<- *model.WebResponse, errCh chan<- error)
	GetDataMerchantId(ctx *fiber.Ctx, request *request.Merchant, ch chan<- *model.WebResponse, errCh chan<- error)
}
