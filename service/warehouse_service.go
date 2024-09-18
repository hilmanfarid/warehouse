package service

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"

	"github.com/go-playground/validator/v10"
)

type WarehouseService interface {
	Create(warehouse request.CreateWarehouseRequest) (response.WarehouseResponse, error)
	Update(warehouse request.UpdateWarehouseRequest) (response.WarehouseResponse, error)
	Delete(warehouse uint32)
	FindById(warehouse uint32) (response.WarehouseResponse, error)
	FindAll() []response.WarehouseResponse
}

type warehouseServiceImpl struct {
	WarehouseRepository repository.WarehouseRepository
	Validate            *validator.Validate
}

func NewWarehouseServiceImpl(warehouseRepository repository.WarehouseRepository, validate *validator.Validate) WarehouseService {
	return &warehouseServiceImpl{
		WarehouseRepository: warehouseRepository,
		Validate:            validate,
	}
}

// Create implements WarehouseService
func (t *warehouseServiceImpl) Create(warehouse request.CreateWarehouseRequest) (response.WarehouseResponse, error) {
	err := t.Validate.Struct(warehouse)
	if err != nil {
		return response.WarehouseResponse{}, err
	}
	params := model.Warehouse{
		Name:   warehouse.Name,
		ShopID: warehouse.ShopID,
		Status: warehouse.Status,
	}
	result, err := t.WarehouseRepository.Save(params)
	if err != nil {
		return response.WarehouseResponse{}, err
	}
	return result.ToResponse(), nil
}

// Delete implements WarehouseService
func (t *warehouseServiceImpl) Delete(warehouseId uint32) {
	t.WarehouseRepository.Delete(warehouseId)
}

// FindAll implements WarehouseService
func (t *warehouseServiceImpl) FindAll() []response.WarehouseResponse {
	result, _ := t.WarehouseRepository.FindAll()

	var warehouse []response.WarehouseResponse
	for _, value := range result {
		warehouse = append(warehouse, value.ToResponse())
	}

	return warehouse
}

// FindById implements WarehouseService
func (t *warehouseServiceImpl) FindById(warehouseId uint32) (response.WarehouseResponse, error) {
	warehouseData, err := t.WarehouseRepository.FindById(warehouseId)
	if err != nil {
		return response.WarehouseResponse{}, err
	}

	return warehouseData.ToResponse(), nil
}

// Update implements WarehouseService
func (t *warehouseServiceImpl) Update(warehouse request.UpdateWarehouseRequest) (response.WarehouseResponse, error) {
	updatedObj, err := t.WarehouseRepository.FindById(warehouse.ID)
	if err != nil {
		return response.WarehouseResponse{}, app_errors.NewMySQLNotFound(err)
	}

	if warehouse.Name != "" {
		updatedObj.Name = warehouse.Name
	}
	if warehouse.Status != 0 {
		updatedObj.Status = warehouse.Status
	}
	if warehouse.ShopID != 0 {
		updatedObj.ShopID = warehouse.ShopID
	}

	result, err := t.WarehouseRepository.Update(updatedObj)
	if err != nil {
		return response.WarehouseResponse{}, nil
	}

	return result.ToResponse(), nil
}
