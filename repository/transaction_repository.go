package repository

import (
	"github.com/gofiber/fiber/v2"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/model/response"
)

type TransactionRepository interface {
	TransactionNotification(ctx *fiber.Ctx, r *request.Transaksi, ch chan error)
	CheckStatus(ctx *fiber.Ctx, r *request.TransaksiStatus, ch chan<- *response.TransactionStatusResponse, errCh chan<- error)
	TransactionList(ctx *fiber.Ctx, ch chan *model.WebResponse, errCh chan error)
	GetBilNumber(ctx *fiber.Ctx, generatedCode string, ch1 chan<- *response.GenerateCodeResponse, errCh chan<- error)
}
