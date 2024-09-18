package service

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"

	"github.com/go-playground/validator/v10"
)

type InventoryService interface {
	Create(inventory request.CreateInventoryRequest) (response.InventoryResponse, error)
	Update(inventory request.UpdateInventoryRequest) (response.InventoryResponse, error)
	Delete(inventory uint32)
	FindById(inventory uint32) (response.InventoryResponse, error)
	FindAll() []response.InventoryResponse
}

type inventoryServiceImpl struct {
	InventoryRepository repository.InventoryRepository
	Validate            *validator.Validate
}

func NewInventoryServiceImpl(inventoryRepository repository.InventoryRepository, validate *validator.Validate) InventoryService {
	return &inventoryServiceImpl{
		InventoryRepository: inventoryRepository,
		Validate:            validate,
	}
}

// Create implements InventoryService
func (t *inventoryServiceImpl) Create(inventory request.CreateInventoryRequest) (response.InventoryResponse, error) {
	err := t.Validate.Struct(inventory)
	if err != nil {
		return response.InventoryResponse{}, err
	}
	params := model.Inventory{
		ProductID:   inventory.ProductID,
		WarehouseID: inventory.WarehouseID,
		Quantity:    inventory.Quantity,
		Status:      inventory.Status,
	}
	result, err := t.InventoryRepository.Save(params)
	if err != nil {
		return response.InventoryResponse{}, err
	}
	return result.ToResponse(), nil
}

// Delete implements InventoryService
func (t *inventoryServiceImpl) Delete(inventoryId uint32) {
	t.InventoryRepository.Delete(inventoryId)
}

// FindAll implements InventoryService
func (t *inventoryServiceImpl) FindAll() []response.InventoryResponse {
	result, _ := t.InventoryRepository.FindAll()

	var inventory []response.InventoryResponse
	for _, value := range result {
		inventory = append(inventory, value.ToResponse())
	}

	return inventory
}

// FindById implements InventoryService
func (t *inventoryServiceImpl) FindById(inventoryId uint32) (response.InventoryResponse, error) {
	inventoryData, err := t.InventoryRepository.FindById(inventoryId)
	if err != nil {
		return response.InventoryResponse{}, err
	}

	return inventoryData.ToResponse(), nil
}

// Update implements InventoryService
func (t *inventoryServiceImpl) Update(inventory request.UpdateInventoryRequest) (response.InventoryResponse, error) {
	updatedObj, err := t.InventoryRepository.FindById(inventory.ID)
	if err != nil {
		return response.InventoryResponse{}, app_errors.NewMySQLNotFound(err)
	}
	if inventory.ProductID != 0 {
		updatedObj.ProductID = inventory.ProductID
	}
	if inventory.Status != 0 {
		updatedObj.Status = inventory.Status
	}
	if inventory.WarehouseID != 0 {
		updatedObj.WarehouseID = inventory.WarehouseID
	}
	if inventory.Quantity != 0 {
		updatedObj.Quantity = inventory.Quantity
	}

	result, err := t.InventoryRepository.Update(updatedObj)
	if err != nil {
		return response.InventoryResponse{}, nil
	}

	return result.ToResponse(), nil
}
