package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"spe_test/middleware"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/service"
	"spe_test/utils/logging"
)

type MerchantController struct {
	MerchantService service.MerchantService
}

func NewMerchantController(companyEmployeeService *service.MerchantService) MerchantController {
	return MerchantController{MerchantService: *companyEmployeeService}
}

func (controller *MerchantController) Route(app *fiber.App) {
	app.Post("/merchant/create", controller.CreateMerchant)
	app.Get("/merchant/list", middleware.AuthMiddleware, controller.ListMerchant)
	app.Get("/merchant/get-data/:id", controller.GetDataMerchantId)

}

func (controller *MerchantController) CreateMerchant(ctx *fiber.Ctx) error {
	var req request.Merchant
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
		response, err := controller.MerchantService.CreateMerchant(ctx, &req)
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

func (controller *MerchantController) ListMerchant(ctx *fiber.Ctx) error {
	ch := make(chan *model.WebResponse)
	errorCh := make(chan error)
	go func() {
		response, err := controller.MerchantService.ListMerchant(ctx)
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

func (controller *MerchantController) GetDataMerchantId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(400).JSON(model.WebResponse{
			Code:    "400",
			Message: "Missing id parameter",
		})
	}

	ch := make(chan *model.WebResponse)
	errCh := make(chan error)

	go func() {
		defer close(ch)
		defer close(errCh)

		response, err := controller.MerchantService.GetDataMerchantId(ctx, id)
		if err != nil {
			errCh <- err
			return
		}
		ch <- response
	}()

	select {
	case response := <-ch:
		return ctx.Status(200).JSON(response)
	case err := <-errCh:
		return ctx.Status(404).JSON(model.WebResponse{
			Code:    "404",
			Message: err.Error(),
		})
	}
}
