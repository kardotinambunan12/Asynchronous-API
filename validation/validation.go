package validation

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"spe_test/model/request"
)

func DataTransactionValidate(ctx *fiber.Ctx, params request.Transaksi) error {
	err := validation.ValidateStruct(&params,
		validation.Field(&params.MerchantID, validation.Required),
		validation.Field(&params.RequestId, validation.Required),
		validation.Field(&params.RRN, validation.Required),
		validation.Field(&params.CustomerPAN, validation.Required),
		validation.Field(&params.TransactionDatetime, validation.Required),
		validation.Field(&params.MerchantCity, validation.Required),
		validation.Field(&params.MerchantName, validation.Required),
		validation.Field(&params.Amount, validation.Required),
		//validation.Field(&params.BillNumber, validation.Required),
		validation.Field(&params.CurrencyCode, validation.Required),
		validation.Field(&params.PetugasRekam, validation.Required),
		validation.Field(&params.PaymentDescription, validation.Required),
		validation.Field(&params.PaymentStatus, validation.Required),
	)

	if err != nil {
		message := err.Error()
		return errors.New(message)
	}
	return nil
}

func DataTransactionStatusValidate(ctx *fiber.Ctx, params request.TransaksiStatus) error {
	err := validation.ValidateStruct(&params,

		validation.Field(&params.BillNumber, validation.Required),
		validation.Field(&params.RequestId, validation.Required),
	)

	if err != nil {
		message := err.Error()
		return errors.New(message)
	}
	return nil
}
