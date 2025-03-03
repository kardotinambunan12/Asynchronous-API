package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"spe_test/config"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/model/response"
)

func NewCustomerRepository() CustomerRepository {
	return &customerRepositoryImpl{}
}

type customerRepositoryImpl struct{}

func (repository *customerRepositoryImpl) InsertCustomer(ctx *fiber.Ctx, request *request.CustomerRequest, ch chan<- error) {
	db := config.NewDB()
	defer db.Close()

	sql := `INSERT INTO spe_test.customer(customer_id, customer_pan, customer_name, tgl_rekam, petugas_rekam) VALUES(?, ?, ?, sysdate(), ?)`
	result, err := db.Exec(sql, request.CostumerId, request.CustomerPAN, request.CustomerName, request.PetugasRekam)

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
func (repository *customerRepositoryImpl) GetDataCustomer(ctx *fiber.Ctx, ch chan<- *model.WebResponse, errCh chan<- error) {
	db := config.NewDB()

	sql := `select customer_id, customer_pan, customer_name, tgl_rekam, petugas_rekam from spe_test.customer`
	rows, err := db.Query(sql)
	if err != nil {
		errCh <- err
		return
	}
	defer rows.Close()

	customers := []response.DataCustomerResponse{}
	for rows.Next() {
		customer := response.DataCustomerResponse{}
		if err := rows.Scan(&customer.CostumerId, &customer.CustomerPAN, &customer.CustomerName, &customer.PetugasRekam, &customer.TglRekam); err != nil {
			errCh <- err
			return
		}
		customers = append(customers, customer)
	}

	resp := &model.WebResponse{
		Code:    "200",
		Message: "Success",
		Data:    customers,
	}
	ch <- resp
}

func (repository *customerRepositoryImpl) GetDataCustomerbyId(ctx *fiber.Ctx, request *request.CustomerRequest, ch chan<- *model.WebResponse, errCh chan<- error) {
	db := config.NewDB()
	defer db.Close() // Pastikan koneksi database ditutup

	sql := `SELECT customer_id, customer_pan, customer_name, tgl_rekam, petugas_rekam FROM spe_test.customer WHERE customer_pan = ?`
	rows, err := db.Query(sql, request.CustomerPAN) // Hilangkan `&`
	if err != nil {
		errCh <- err
		return
	}
	defer rows.Close()

	customer := response.DataCustomerResponse{}
	if rows.Next() {
		if err := rows.Scan(&customer.CostumerId, &customer.CustomerPAN, &customer.CustomerName, &customer.PetugasRekam, &customer.TglRekam); err != nil {
			errCh <- err
			return
		}

		resp := &model.WebResponse{
			Code:    "200",
			Message: "Success",
			Data:    customer,
		}
		ch <- resp
	} else {
		errCh <- errors.New("Customer not found")
	}
}
