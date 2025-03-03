package service

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"spe_test/config"
	errorhandler "spe_test/error_handler"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/model/response"
	"spe_test/repository"
	"spe_test/utils"
	"spe_test/validation"
	"strconv"
)

func NewTransactionService(transactionRepository *repository.TransactionRepository, config config.Config) TransationService {
	return &transactionServiceImpl{
		TransactionRepository: *transactionRepository,
		Configuration:         config,
	}
}

type transactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	Configuration         config.Config
}

func (service *transactionServiceImpl) TransactionList(ctx *fiber.Ctx) (*model.WebResponse, error) {
	ch := make(chan *model.WebResponse)
	errCh := make(chan error)
	go service.TransactionRepository.TransactionList(ctx, ch, errCh)

	select {
	case res := <-ch:
		return res, nil
	case err := <-errCh:
		return nil, errors.New("Failed to get data Transaction: " + err.Error())
	}
}

func (service *transactionServiceImpl) CheckStatus(ctx *fiber.Ctx, r *request.TransaksiStatus) (*response.TransactionStatusResponse, error) {
	errValidation := validation.DataTransactionStatusValidate(ctx, *r)
	if errValidation != nil {
		return nil, errValidation
	}

	signaturePlain := r.BillNumber
	apiKey := utils.GenerateSignature(signaturePlain, service.Configuration.Get("SIGNATURE_KEY"))
	log.Println("apiKey", apiKey)

	signatureRemote := ctx.Get("X-Signature")
	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(signatureRemote)) != 1 {
		message := "ERROR_CREDENTIAL_NOT_VALID"
		return nil, errorhandler.GeneralError{
			Message: message,
		}
	}

	ch := make(chan *response.TransactionStatusResponse)
	errCh := make(chan error)

	go service.TransactionRepository.CheckStatus(ctx, r, ch, errCh)

	select {
	case res := <-ch:
		amountFloat, err := strconv.ParseFloat(res.Amount, 64)
		if err != nil {
			return nil, errors.New("Failed to parse amount: ")
		}
		formattedAmount := fmt.Sprintf("%.2f", amountFloat)

		res = &response.TransactionStatusResponse{
			Code:                "00",
			Message:             "Success",
			RequestID:           res.RequestID,
			CustomerPAN:         res.CustomerPAN,
			Amount:              formattedAmount,
			TransactionDatetime: res.TransactionDatetime,
			RRN:                 res.RRN,
			BillNumber:          res.BillNumber,
			CustomerName:        res.CustomerName,
			MerchantID:          res.MerchantID,
			MerchantName:        res.MerchantName,
			MerchantCity:        res.MerchantCity,
			CurrencyCode:        res.CurrencyCode,
			PaymentStatus:       res.PaymentStatus,
			PaymentDescription:  res.PaymentDescription,
		}
		return res, nil
	case err := <-errCh:
		return nil, errors.New("data merchant not found: " + err.Error())
	}

}

func (service *transactionServiceImpl) TransactionNotification(ctx *fiber.Ctx, r *request.Transaksi) (*model.WebResponse, error) {
	errValidation := validation.DataTransactionValidate(ctx, *r)
	if errValidation != nil {
		return nil, errValidation
	}

	signaturePlain := r.RequestId + ":" + r.RRN + ":" + r.MerchantID
	apiKey := utils.GenerateSignature(signaturePlain, service.Configuration.Get("SIGNATURE_KEY"))
	log.Println("apiKey", apiKey)
	signatureRemote := ctx.Get("X-Signature")
	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(signatureRemote)) != 1 {
		message := "ERROR_CREDENTIAL_NOT_VALID"
		return nil, errorhandler.GeneralError{
			Message: message,
		}
	}

	ch := make(chan error)
	ch1 := make(chan *response.GenerateCodeResponse)

	//id, _ := uuid.NewRandom()
	//trxId := strings.ReplaceAll(id.String(), "-", "")
	//
	//r.RequestId = trxId

	counter := utils.GenerateCounter()
	generatedCode := fmt.Sprintf("INV%08d", counter)

	go service.TransactionRepository.GetBilNumber(ctx, generatedCode, ch1, ch)

	generatedCodeResponse := <-ch1

	if generatedCodeResponse.BillNumber == "" {
		r.BillNumber = generatedCode
	} else {
		kodeKategori := generatedCodeResponse.BillNumber[3:]

		nomor, err := strconv.Atoi(kodeKategori)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		counter1 := utils.GenerateLastCounter(nomor)
		generatedLastCode := fmt.Sprintf("INV%08d", counter1)

		r.BillNumber = generatedLastCode

	}

	go service.TransactionRepository.TransactionNotification(ctx, r, ch)

	err := <-ch
	if err != nil {
		return nil, errors.New("Failed to insert transaction")
	}

	response := &model.WebResponse{
		Code:    "00",
		Message: "Success",
	}
	return response, nil
}
