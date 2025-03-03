package service

import (
	"github.com/gofiber/fiber/v2"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/model/response"
)

type TransationService interface {
	TransactionNotification(ctx *fiber.Ctx, r *request.Transaksi) (*model.WebResponse, error)
	CheckStatus(ctx *fiber.Ctx, r *request.TransaksiStatus) (*response.TransactionStatusResponse, error)
	TransactionList(ctx *fiber.Ctx) (*model.WebResponse, error)
}
