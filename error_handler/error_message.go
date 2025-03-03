package errorhandler

import (
	"github.com/gofiber/fiber/v2"
	"spe_test/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, databaseError := err.(DatabaseError)
	if databaseError {
		return ctx.Status(500).JSON(model.GeneralResponse{
			StatusCode: 500,
			Message:    "Terjadi kesalahan, silahkan ulangi beberapa saat lagi",
		})
	}

	_, dataNotFoundError := err.(DataNotFoundError)
	if dataNotFoundError {
		return ctx.Status(404).JSON(model.GeneralResponse{
			StatusCode: 404,
			Message:    err.Error(),
		})
	}

	_, generalError := err.(GeneralError)
	if generalError {
		return ctx.Status(400).JSON(model.GeneralResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	return ctx.Status(400).JSON(model.GeneralResponse{
		StatusCode: 400,
		Message:    err.Error(),
	})
}
