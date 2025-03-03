package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"spe_test/config"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/model/response"
)

func NewMerchantpository() MerchantRepository {
	return &merchantRepositoryImpl{}
}

type merchantRepositoryImpl struct{}

func (m *merchantRepositoryImpl) CreateMerchant(ctx *fiber.Ctx, request *request.Merchant, ch chan<- error) {
	db := config.NewDB()
	defer db.Close()

	sql := `INSERT INTO spe_test.merchant(merchant_id, merchant_name, merchant_city, tgl_rekam, petugas_rekam) VALUES(?, ?, ?, sysdate(), ?)`
	result, err := db.Exec(sql, request.MerchantID, request.MerchantName, request.MerchantCity, request.PetugasRekam)

	if err != nil {
		ch <- err
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ch <- err
		return
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)

	ch <- nil
}

func (m *merchantRepositoryImpl) ListMerchant(ctx *fiber.Ctx, ch chan<- *model.WebResponse, errCh chan<- error) {
	db := config.NewDB()

	sql := `select merchant_id, merchant_name, merchant_city, tgl_rekam, petugas_rekam from spe_test.merchant`
	rows, err := db.Query(sql)
	if err != nil {
		errCh <- err
		return
	}
	defer rows.Close()

	merchants := []response.DataMerchantResponse{}
	for rows.Next() {
		merchant := response.DataMerchantResponse{}
		if err := rows.Scan(&merchant.MerchantID, &merchant.MerchantName, &merchant.MerchantCity, &merchant.PetugasRekam, &merchant.TglRekam); err != nil {
			errCh <- err
			return
		}
		merchants = append(merchants, merchant)
	}

	resp := &model.WebResponse{
		Code:    "200",
		Message: "Success",
		Data:    merchants,
	}
	ch <- resp
}

func (repository *merchantRepositoryImpl) GetDataMerchantId(ctx *fiber.Ctx, request *request.Merchant, ch chan<- *model.WebResponse, errCh chan<- error) {
	db := config.NewDB()

	sql := `select merchant_id, merchant_name, merchant_city, tgl_rekam, petugas_rekam from spe_test.merchant where merchant_id = ?`
	rows, err := db.Query(sql, request.MerchantID)
	if err != nil {
		errCh <- err
		return
	}
	defer rows.Close()

	merchant := response.DataMerchantResponse{}
	for rows.Next() {
		if err := rows.Scan(&merchant.MerchantID, &merchant.MerchantName, &merchant.MerchantCity, &merchant.PetugasRekam, &merchant.TglRekam); err != nil {
			errCh <- err
			return
		}
	}
	if merchant.MerchantID == "" {
		errCh <- fmt.Errorf("Data not found")
		return
	}

	resp := &model.WebResponse{
		Code:    "200",
		Message: "Success",
		Data:    merchant,
	}
	ch <- resp
}
