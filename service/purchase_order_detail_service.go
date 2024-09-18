package service

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"

	"github.com/go-playground/validator/v10"
)

type PurchaseOrderDetailService interface {
	ProcessDetail(purchaseOrderDetailId uint32, state model.OrderDetailStatus) (response.PurchaseOrderDetailResponse, error)
	Update(purchaseOrderDetail request.UpdatePurchaseOrderDetailsRequest) (response.PurchaseOrderDetailResponse, error)
	Delete(purchaseOrderDetail uint32)
	FindById(purchaseOrderDetail uint32) (response.PurchaseOrderDetailResponse, error)
	FindByPurchaseOrderId(purchaseOrderDetail uint32) ([]response.PurchaseOrderDetailResponse, error)
	FindAll() []response.PurchaseOrderDetailResponse
}

type purchaseOrderDetailServiceImpl struct {
	PurchaseOrderDetailRepository repository.PurchaseOrderDetailRepository
	PurchaseOrderRepository       repository.PurchaseOrderRepository
	InventoryService              InventoryService
	Validate                      *validator.Validate
}

func NewPurchaseOrderDetailServiceImpl(purchaseOrderDetailRepository repository.PurchaseOrderDetailRepository, purchaseOrderRepository repository.PurchaseOrderRepository, inventoryService InventoryService, validate *validator.Validate) PurchaseOrderDetailService {
	return &purchaseOrderDetailServiceImpl{
		InventoryService:              inventoryService,
		PurchaseOrderDetailRepository: purchaseOrderDetailRepository,
		PurchaseOrderRepository:       purchaseOrderRepository,
		Validate:                      validate,
	}
}

// Delete implements PurchaseOrderDetailService
func (t *purchaseOrderDetailServiceImpl) Delete(purchaseOrderDetailId uint32) {
	t.PurchaseOrderDetailRepository.Delete(purchaseOrderDetailId)
}

// FindAll implements PurchaseOrderDetailService
func (t *purchaseOrderDetailServiceImpl) FindAll() []response.PurchaseOrderDetailResponse {
	result, _ := t.PurchaseOrderDetailRepository.FindAll()

	var purchaseOrderDetail []response.PurchaseOrderDetailResponse
	for _, value := range result {
		purchaseOrderDetail = append(purchaseOrderDetail, value.ToResponse())
	}

	return purchaseOrderDetail
}

// FindById implements PurchaseOrderDetailService
func (t *purchaseOrderDetailServiceImpl) FindById(purchaseOrderDetailId uint32) (response.PurchaseOrderDetailResponse, error) {
	purchaseOrderDetailData, err := t.PurchaseOrderDetailRepository.FindById(purchaseOrderDetailId)
	if err != nil {
		return response.PurchaseOrderDetailResponse{}, err
	}

	return purchaseOrderDetailData.ToResponse(), nil
}

func (t *purchaseOrderDetailServiceImpl) ProcessDetail(purchaseOrderDetailId uint32, state model.OrderDetailStatus) (response.PurchaseOrderDetailResponse, error) {
	purchaseOrderDetailData, err := t.PurchaseOrderDetailRepository.FindById(purchaseOrderDetailId)
	if err != nil {
		return response.PurchaseOrderDetailResponse{}, err
	}

	if purchaseOrderDetailData.Status >= model.OrderDetailStatusFailed {
		return purchaseOrderDetailData.ToResponse(), app_errors.NewInvalidStateError(nil)
	}

	purchaseOrderData, err := t.PurchaseOrderRepository.FindById(purchaseOrderDetailData.PurchaseOrderID)
	if err != nil {
		return purchaseOrderDetailData.ToResponse(), err
	}

	detail, err := t.PurchaseOrderDetailRepository.PurchaseStock(purchaseOrderData, purchaseOrderDetailData, state)
	if err != nil {
		return purchaseOrderDetailData.ToResponse(), err
	}
	return detail.ToResponse(), nil
}

func (t *purchaseOrderDetailServiceImpl) FindByPurchaseOrderId(purchaseOrderID uint32) ([]response.PurchaseOrderDetailResponse, error) {
	purchaseOrderDetailData, err := t.PurchaseOrderDetailRepository.FindByPurchaseOrderID(purchaseOrderID)
	if err != nil {
		return []response.PurchaseOrderDetailResponse{}, err
	}

	var purchaseOrderDetailResponse []response.PurchaseOrderDetailResponse
	for _, detail := range purchaseOrderDetailData {
		purchaseOrderDetailResponse = append(purchaseOrderDetailResponse, detail.ToResponse())
	}
	return purchaseOrderDetailResponse, nil
}

// Update implements PurchaseOrderDetailService
func (t *purchaseOrderDetailServiceImpl) Update(purchaseOrderDetail request.UpdatePurchaseOrderDetailsRequest) (response.PurchaseOrderDetailResponse, error) {

	panic("implement mme")
}
