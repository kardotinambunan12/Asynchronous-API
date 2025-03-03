package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"spe_test/model"
	"spe_test/model/request"
	"spe_test/repository"
)

func NewMerchantService(merchantRepository *repository.MerchantRepository) MerchantService {
	return &merchantServiceImpl{
		MerchantRepository: *merchantRepository,
	}
}

type merchantServiceImpl struct {
	MerchantRepository repository.MerchantRepository
}

func (service *merchantServiceImpl) ListMerchant(ctx *fiber.Ctx) (*model.WebResponse, error) {

	ch := make(chan *model.WebResponse)
	errCh := make(chan error)
	go service.MerchantRepository.ListMerchant(ctx, ch, errCh)

	select {
	case res := <-ch:
		return res, nil
	case err := <-errCh:
		return nil, errors.New("Failed to get data customer: " + err.Error())
	}
}

func (service *merchantServiceImpl) CreateMerchant(ctx *fiber.Ctx, r *request.Merchant) (*model.WebResponse, error) {

	ch := make(chan error)

	id, _ := uuid.NewRandom()

	req := &request.Merchant{
		MerchantID:   id.String(),
		MerchantName: r.MerchantName,
		MerchantCity: r.MerchantCity,
		PetugasRekam: r.PetugasRekam,
		TglRekam:     r.TglRekam,
	}

	go service.MerchantRepository.CreateMerchant(ctx, req, ch)

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

func (service *merchantServiceImpl) GetDataMerchantId(ctx *fiber.Ctx, merchantId string) (*model.WebResponse, error) {
	ch := make(chan *model.WebResponse)
	errCh := make(chan error)

	r := &request.Merchant{
		MerchantID: merchantId,
	}

	go service.MerchantRepository.GetDataMerchantId(ctx, r, ch, errCh)

	select {
	case res := <-ch:
		return res, nil
	case err := <-errCh:
		return nil, errors.New("data merchant not found: " + err.Error())
	}
}
