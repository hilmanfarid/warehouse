package service

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"

	"github.com/go-playground/validator/v10"
)

type ShopService interface {
	Create(shop request.CreateShopRequest) (response.ShopResponse, error)
	Update(shop request.UpdateShopRequest) (response.ShopResponse, error)
	Delete(shop uint32)
	FindById(shop uint32) (response.ShopResponse, error)
	FindAll() []response.ShopResponse
}

type shopServiceImpl struct {
	ShopRepository repository.ShopRepository
	Validate       *validator.Validate
}

func NewShopServiceImpl(shopRepository repository.ShopRepository, validate *validator.Validate) ShopService {
	return &shopServiceImpl{
		ShopRepository: shopRepository,
		Validate:       validate,
	}
}

// Create implements ShopService
func (t *shopServiceImpl) Create(shop request.CreateShopRequest) (response.ShopResponse, error) {
	err := t.Validate.Struct(shop)
	if err != nil {
		return response.ShopResponse{}, err
	}
	params := model.Shop{
		Name:   shop.Name,
		Status: shop.Status,
	}
	result, err := t.ShopRepository.Save(params)
	if err != nil {
		return response.ShopResponse{}, err
	}
	return result.ToResponse(), nil
}

// Delete implements ShopService
func (t *shopServiceImpl) Delete(shopId uint32) {
	t.ShopRepository.Delete(shopId)
}

// FindAll implements ShopService
func (t *shopServiceImpl) FindAll() []response.ShopResponse {
	result, _ := t.ShopRepository.FindAll()

	var shop []response.ShopResponse
	for _, value := range result {
		shop = append(shop, value.ToResponse())
	}

	return shop
}

// FindById implements ShopService
func (t *shopServiceImpl) FindById(shopId uint32) (response.ShopResponse, error) {
	shopData, err := t.ShopRepository.FindById(shopId)
	if err != nil {
		return response.ShopResponse{}, err
	}

	return shopData.ToResponse(), nil
}

// Update implements ShopService
func (t *shopServiceImpl) Update(shop request.UpdateShopRequest) (response.ShopResponse, error) {
	updatedObj, err := t.ShopRepository.FindById(shop.ID)
	if err != nil {
		return response.ShopResponse{}, app_errors.NewMySQLNotFound(err)
	}

	if shop.Name != "" {
		updatedObj.Name = shop.Name
	}
	if shop.Status != 0 {
		updatedObj.Status = shop.Status
	}

	result, err := t.ShopRepository.Update(updatedObj)
	if err != nil {
		return response.ShopResponse{}, nil
	}

	return result.ToResponse(), nil
}
