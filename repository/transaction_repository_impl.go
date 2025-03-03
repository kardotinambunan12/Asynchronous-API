package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"spe_test/config"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/model/response"
)

func NewTransactionRepository() TransactionRepository {
	return &transactionRepositoryImpl{}
}

type transactionRepositoryImpl struct{}

func (repository *transactionRepositoryImpl) GetBilNumber(ctx *fiber.Ctx, generatedCode string, ch chan<- *response.GenerateCodeResponse, errCh chan<- error) {
	db := config.NewDB()

	sql := `SELECT bill_number
					FROM spe_test.transactions
					ORDER BY bill_number DESC
					LIMIT 1`
	rows, err := db.Query(sql)
	if err != nil {
		errCh <- err
		return
	}
	defer rows.Close()
	codeBillNumber := &response.GenerateCodeResponse{}
	for rows.Next() {

		if err := rows.Scan(
			&codeBillNumber.BillNumber); err != nil {
			errCh <- err
			return
		}

	}

	resp := &response.GenerateCodeResponse{
		BillNumber: codeBillNumber.BillNumber,
	}
	ch <- resp
}

func (repository *transactionRepositoryImpl) TransactionList(ctx *fiber.Ctx, ch chan *model.WebResponse, errCh chan error) {
	db := config.NewDB()

	sql := `select request_id, customer_pan, amount, 
					   transaction_datetime, rrn, bill_number, 
					   customer_name, merchant_id, merchant_name, 
					   merchant_city, currency_code, payment_status, 
					   payment_description
				from spe_test.transactions`
	rows, err := db.Query(sql)
	if err != nil {
		errCh <- err
		return
	}
	defer rows.Close()

	trxs := []response.TransactionStatusResponse{}
	for rows.Next() {
		trx := response.TransactionStatusResponse{}

		if err := rows.Scan(
			&trx.RequestID,
			&trx.CustomerPAN,
			&trx.Amount,
			&trx.TransactionDatetime,
			&trx.RRN,
			&trx.BillNumber,
			&trx.CustomerName,
			&trx.MerchantID,
			&trx.MerchantName,
			&trx.MerchantCity,
			&trx.CurrencyCode,
			&trx.PaymentStatus,
			&trx.PaymentDescription); err != nil {
			errCh <- err
			return
		}
		trxs = append(trxs, trx)
	}

	resp := &model.WebResponse{
		Code:    "200",
		Message: "Success",
		Data:    trxs,
	}
	ch <- resp
}

func (repository *transactionRepositoryImpl) CheckStatus(ctx *fiber.Ctx, r *request.TransaksiStatus, ch chan<- *response.TransactionStatusResponse, errCh chan<- error) {
	db := config.NewDB()

	sql := `select request_id, customer_pan, amount, 
					   transaction_datetime, rrn, bill_number, 
					   customer_name, merchant_id, merchant_name, 
					   merchant_city, currency_code, payment_status, 
					   payment_description 
				from spe_test.transactions where request_id = ?
				and bill_number = ?`
	rows, err := db.Query(sql, r.RequestId, r.BillNumber)
	if err != nil {
		errCh <- err
		return
	}
	defer rows.Close()

	trx := &response.TransactionStatusResponse{}
	for rows.Next() {
		if err := rows.Scan(
			&trx.RequestID,
			&trx.CustomerPAN,
			&trx.Amount,
			&trx.TransactionDatetime,
			&trx.RRN,
			&trx.BillNumber,
			&trx.CustomerName,
			&trx.MerchantID,
			&trx.MerchantName,
			&trx.MerchantCity,
			&trx.CurrencyCode,
			&trx.PaymentStatus,
			&trx.PaymentDescription); err != nil {
			errCh <- err
			return
		}
	}
	if trx.MerchantID == "" {
		errCh <- fmt.Errorf("Data not found")
		return
	}

	ch <- trx
}

func (repository *transactionRepositoryImpl) TransactionNotification(ctx *fiber.Ctx, r *request.Transaksi, ch chan error) {
	db := config.NewDB()

	sql := `INSERT INTO spe_test.transactions 
		(request_id, customer_pan, amount, transaction_datetime, rrn, bill_number, customer_name, 
		 merchant_id, merchant_name, merchant_city, currency_code, payment_status, payment_description, tgl_rekam, petugas_rekam) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, sysdate(), ?)`

	_, err := db.Exec(sql, r.RequestId, r.CustomerPAN, r.Amount, r.TransactionDatetime, r.RRN,
		r.BillNumber, r.CustomerName, r.MerchantID, r.MerchantName, r.MerchantCity,
		r.CurrencyCode, r.PaymentStatus, r.PaymentDescription, r.PetugasRekam)
	if err != nil {
		ch <- err
		return
	}

	defer db.Close()
	ch <- nil
}
