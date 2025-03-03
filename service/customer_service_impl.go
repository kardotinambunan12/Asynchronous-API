package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/repository"
)

func NewCustomerService(customerRepository *repository.CustomerRepository) CustomerService {
	return &customerImpl{
		CustomerRepository: *customerRepository,
	}
}

type customerImpl struct {
	CustomerRepository repository.CustomerRepository
}

func (service *customerImpl) InsertCustomer(ctx *fiber.Ctx, r *request.CustomerRequest) (*model.WebResponse, error) {
	ch := make(chan error)

	id, _ := uuid.NewRandom()

	req := &request.CustomerRequest{
		CostumerId:   id.String(),
		CustomerPAN:  r.CustomerPAN,
		CustomerName: r.CustomerName,
		PetugasRekam: r.PetugasRekam,
	}

	go service.CustomerRepository.InsertCustomer(ctx, req, ch)

	err := <-ch
	if err != nil {
		return nil, errors.New("Failed to insert customer")
	}

	response := &model.WebResponse{
		Code:    "200",
		Message: "Success",
	}
	return response, nil
}

func (service *customerImpl) GetDataCustomer(ctx *fiber.Ctx) (*model.WebResponse, error) {
	ch := make(chan *model.WebResponse)
	errCh := make(chan error)
	go service.CustomerRepository.GetDataCustomer(ctx, ch, errCh)

	select {
	case res := <-ch:
		return res, nil
	case err := <-errCh:
		return nil, errors.New("Failed to get data customer: " + err.Error())
	}
}

func (service *customerImpl) GetDataCustomerbyId(ctx *fiber.Ctx, customerPan string) (*model.WebResponse, error) {
	ch := make(chan *model.WebResponse)
	errCh := make(chan error)

	r := &request.CustomerRequest{
		CustomerPAN: customerPan,
	}

	go service.CustomerRepository.GetDataCustomerbyId(ctx, r, ch, errCh)

	select {
	case res := <-ch:
		return res, nil
	case err := <-errCh:
		return nil, errors.New("Failed to get customer data: " + err.Error())
	}
}
