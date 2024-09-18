package service

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"

	"github.com/go-playground/validator/v10"
)

type TransferOrderService interface {
	TransferOrder(transferOrder request.CreateTransferOrderRequest) (response.TransferOrderResponse, error)
	FindById(TransferOrderId uint32) (response.TransferOrderResponse, error)
	ProcessTransfer(transferOrderID uint32) (response.TransferOrderResponse, error)
}

type transferOrderServiceImpl struct {
	ProductService          ProductService
	InventoryRepository     repository.InventoryRepository
	TransferOrderRepository repository.TransferOrderRepository
	Validate                *validator.Validate
}

func NewTransferOrderServiceImpl(inventoryRepository repository.InventoryRepository, transferOrderRepository repository.TransferOrderRepository, validate *validator.Validate) TransferOrderService {
	return &transferOrderServiceImpl{
		InventoryRepository:     inventoryRepository,
		TransferOrderRepository: transferOrderRepository,
		Validate:                validate,
	}
}

func (t *transferOrderServiceImpl) TransferOrder(transferOrder request.CreateTransferOrderRequest) (response.TransferOrderResponse, error) {
	err := t.Validate.Struct(transferOrder)
	if err != nil {
		return response.TransferOrderResponse{}, err
	}

	transfer := model.TransferOrder{
		UserID:               transferOrder.UserID,
		ProductID:            transferOrder.ProductID,
		SourceWarehouse:      transferOrder.SourceWarehouse,
		DestinationWarehouse: transferOrder.DestinationWarehouse,
		Status:               transferOrder.Status,
		Quantity:             transferOrder.Quantity,
	}

	result, err := t.TransferOrderRepository.Save(transfer)
	if err != nil {
		return response.TransferOrderResponse{}, err
	}
	return result.ToResponse(), nil
}

func (t *transferOrderServiceImpl) ProcessTransfer(transferOrderID uint32) (response.TransferOrderResponse, error) {
	transferOrderData, err := t.TransferOrderRepository.FindById(transferOrderID)
	if err != nil {
		return response.TransferOrderResponse{}, err
	}
	if transferOrderData.Status > model.TransferOrderStatusCreated {
		return transferOrderData.ToResponse(), app_errors.NewInvalidStateError(nil)
	}

	err = t.InventoryRepository.TransferStock(transferOrderData)
	if err != nil {
		transferOrderData.Status = model.TransferOrderStatusFailed
	} else {
		transferOrderData.Status = model.TransferOrderStatusSucceeded
	}
	updatedData, err := t.TransferOrderRepository.Update(transferOrderData)
	if err != nil {
		return transferOrderData.ToResponse(), err
	}
	return updatedData.ToResponse(), nil
}

func (t *transferOrderServiceImpl) FindById(TransferOrderId uint32) (response.TransferOrderResponse, error) {
	transferOrderData, err := t.TransferOrderRepository.FindById(TransferOrderId)
	if err != nil {
		return response.TransferOrderResponse{}, err
	}
	return transferOrderData.ToResponse(), nil
}
