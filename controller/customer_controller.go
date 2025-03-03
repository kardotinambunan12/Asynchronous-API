package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/service"
	"spe_test/utils/logging"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) CustomerController {
	return CustomerController{CustomerService: *customerService}
}
func (controller *CustomerController) Route(app *fiber.App) {
	app.Post("/customer/insert", controller.InsertCustomer)
	app.Get("/customer/get-data", controller.GetDataCustomer)
	app.Get("/customer/get-data-id/:id", controller.GetDataCustomerbyId)

}

func (controller *CustomerController) InsertCustomer(ctx *fiber.Ctx) error {
	var req request.CustomerRequest
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
		response, err := controller.CustomerService.InsertCustomer(ctx, &req)
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

func (controller *CustomerController) GetDataCustomer(ctx *fiber.Ctx) error {
	ch := make(chan *model.WebResponse)
	errorCh := make(chan error)
	go func() {
		response, err := controller.CustomerService.GetDataCustomer(ctx)
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
func (controller *CustomerController) GetDataCustomerbyId(ctx *fiber.Ctx) error {
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

		response, err := controller.CustomerService.GetDataCustomerbyId(ctx, id)
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
