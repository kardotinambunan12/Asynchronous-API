package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"spe_test/middleware"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/model/response"
	"spe_test/service"
	"spe_test/utils/logging"
)

type TransactionController struct {
	TransactionService service.TransationService
}

func NewTransactionController(companyEmployeeUserService *service.TransationService) TransactionController {
	return TransactionController{TransactionService: *companyEmployeeUserService}
}

func (controller *TransactionController) Route(app *fiber.App) {
	app.Post("/transaction-notification", middleware.AuthMiddleware, controller.TransactionNotification)
	app.Post("/check-status", middleware.AuthMiddleware, controller.CheckStatus)
	app.Post("/transaction-list", middleware.AuthMiddleware, controller.TransactionList)

}

func (controller *TransactionController) TransactionNotification(ctx *fiber.Ctx) error {
	var req request.Transaksi
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	requestId := guuid.New()
	logStart := logging.LogRequest(ctx, string(requestData), requestId.String(), "0000")

	ch := make(chan *model.WebResponse)
	errorCh := make(chan error)
	go func() {
		response, err := controller.TransactionService.TransactionNotification(ctx, &req)
		if err != nil {
			errorCh <- err
			return
		}
		ch <- response
	}()

	select {
	case response := <-ch:
		responseData, err := json.Marshal(response)
		if err != nil {
			return err
		}
		logging.LogResponse(ctx, string(responseData), requestId.String(), "0000", logStart)
		return ctx.Status(200).JSON(response)
	case err := <-errorCh:
		return err
	}
}

func (controller *TransactionController) CheckStatus(ctx *fiber.Ctx) error {
	var req request.TransaksiStatus
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	requestId := guuid.New()
	logStart := logging.LogRequest(ctx, string(requestData), requestId.String(), "0000")

	ch := make(chan *response.TransactionStatusResponse)
	errorCh := make(chan error)
	go func() {
		response, err1 := controller.TransactionService.CheckStatus(ctx, &req)

		if err1 != nil {
			errorCh <- err1
			return
		}
		ch <- response
	}()

	select {
	case response := <-ch:
		responseData, err := json.Marshal(response)
		if err != nil {
			return err
		}
		logging.LogResponse(ctx, string(responseData), requestId.String(), "0000", logStart)

		return ctx.Status(200).JSON(response)
	case err := <-errorCh:
		return err
	}
}

func (controller *TransactionController) TransactionList(ctx *fiber.Ctx) error {
	ch := make(chan *model.WebResponse)
	errorCh := make(chan error)
	go func() {
		response, err := controller.TransactionService.TransactionList(ctx)
		if err != nil {
			errorCh <- err
			return
		}
		ch <- response
	}()

	select {
	case response := <-ch:
		return ctx.Status(200).JSON(response)
	case err := <-errorCh:
		return err
	}
}
