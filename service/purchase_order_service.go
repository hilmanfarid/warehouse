package service

import (
	"fmt"
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type PurchaseOrderService interface {
	CreateOrder(PurchaseOrder request.CreatePurchaseOrderRequest) (response.PurchaseOrderResponse, error)
	FindByIdWithUser(PurchaseOrder uint32, userId uint32) (response.PurchaseOrderResponse, error)
	FindAll() []response.PurchaseOrderResponse
	ProcessOrder(purchaseOrderID uint32) (response.PurchaseOrderResponse, error)
	RefundOrder(purchaseOrderID uint32) (response.PurchaseOrderResponse, error)
	MassRefund() error
}

type purchaseOrderServiceImpl struct {
	ProductService             ProductService
	PurchaseOrderDetailService PurchaseOrderDetailService
	PurchaseOrderRepository    repository.PurchaseOrderRepository
	Validate                   *validator.Validate
}

func NewPurchaseOrderServiceImpl(productService ProductService, PurchaseOrderDetailService PurchaseOrderDetailService, purchaseOrderRepository repository.PurchaseOrderRepository, validate *validator.Validate) PurchaseOrderService {
	return &purchaseOrderServiceImpl{
		ProductService:             productService,
		PurchaseOrderDetailService: PurchaseOrderDetailService,
		PurchaseOrderRepository:    purchaseOrderRepository,
		Validate:                   validate,
	}
}

// CreateOrder create with details
func (t *purchaseOrderServiceImpl) CreateOrder(purchaseOrder request.CreatePurchaseOrderRequest) (response.PurchaseOrderResponse, error) {
	err := t.Validate.Struct(purchaseOrder)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}

	order := model.PurchaseOrder{
		UserID: purchaseOrder.UserID,
		ShopID: purchaseOrder.ShopID,
		Status: purchaseOrder.Status,
	}
	orderDetails := make([]model.PurchaseOrderDetail, 0)
	for _, detail := range purchaseOrder.OrderDetails {
		product, err := t.ProductService.FindById(detail.ProductID)
		if err != nil {
			return response.PurchaseOrderResponse{}, err
		}
		orderDetail := model.PurchaseOrderDetail{
			ProductID:    detail.ProductID,
			Status:       1,
			Quantity:     detail.Quantity,
			PricePerUnit: product.Price,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	result, err := t.PurchaseOrderRepository.CreateOrder(order, orderDetails)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}
	return result, nil
}

func (t *purchaseOrderServiceImpl) RefundOrder(purchaseOrderID uint32) (response.PurchaseOrderResponse, error) {
	log.Info().Msg("Start unstuck order")
	po, err := t.PurchaseOrderRepository.FindById(purchaseOrderID)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}
	if po.Status != model.OrderStatusProcessed {
		return response.PurchaseOrderResponse{}, app_errors.NewInvalidStateError(nil)
	}

	// process the products stock
	details, orderState := t.ProcessRefundStock(po.ID)
	po.Status = orderState

	_, err = t.PurchaseOrderRepository.Update(po)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}
	orderResponse := po.ToResponse()
	orderResponse.OrderDetails = details

	log.Info().Msg("Finish unstuck order")
	return orderResponse, nil
}

func (t *purchaseOrderServiceImpl) ProcessOrder(purchaseOrderID uint32) (response.PurchaseOrderResponse, error) {
	po, err := t.PurchaseOrderRepository.FindById(purchaseOrderID)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}

	// set status processed
	po.Status = model.OrderStatusProcessed
	po.ProcessAt = time.Now()
	_, err = t.PurchaseOrderRepository.Update(po)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}

	// process the products stock
	details, orderState, totalAmount := t.ProcessDetail(po.ID)
	po.Status = orderState
	po.TotalAmount = totalAmount
	_, err = t.PurchaseOrderRepository.Update(po)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}
	orderResponse := po.ToResponse()
	orderResponse.OrderDetails = details
	return orderResponse, nil
}

func (t *purchaseOrderServiceImpl) ProcessDetail(purchaseOrderID uint32) (pds []response.PurchaseOrderDetailResponse, state model.OrderStatus, totalAmount float64) {
	pds, err := t.PurchaseOrderDetailService.FindByPurchaseOrderId(purchaseOrderID)
	if err != nil {
		return pds, model.OrderStatusFailed, 0
	}

	var numOfSuccess int
	for i, detail := range pds {
		totalAmount += detail.PricePerUnit * float64(detail.Quantity)
		pds[i], err = t.PurchaseOrderDetailService.ProcessDetail(detail.ID, model.OrderDetailStatusSucceeded)
		if err != nil {
			continue
		}
		if pds[i].Status == model.OrderDetailStatusSucceeded.String() {
			numOfSuccess++
		}
	}

	var status model.OrderStatus
	if numOfSuccess == 0 {
		status = model.OrderStatusFailed
	} else if numOfSuccess == len(pds) {
		status = model.OrderStatusSucceeded
	} else {
		status = model.OrderStatusProcessed
	}
	return pds, status, totalAmount
}

func (t *purchaseOrderServiceImpl) ProcessRefundStock(purchaseOrderID uint32) (pds []response.PurchaseOrderDetailResponse, state model.OrderStatus) {
	pds, err := t.PurchaseOrderDetailService.FindByPurchaseOrderId(purchaseOrderID)
	if err != nil {
		return pds, model.OrderStatusFailed
	}

	var numOfSuccess int
	for i, detail := range pds {
		if model.OderStatusFromString(detail.Status) >= model.OrderStatusFailed {
			numOfSuccess++
			continue
		}
		pds[i], err = t.PurchaseOrderDetailService.ProcessDetail(detail.ID, model.OrderDetailStatusRefunded)
		if err != nil {
			log.Error().Msg(err.Error())
			continue
		}
		if pds[i].Status == model.OrderDetailStatusRefunded.String() {
			numOfSuccess++
		}
	}

	if numOfSuccess == len(pds) {
		log.Info().Msg(fmt.Sprintf("order id : %d successfully refunded", purchaseOrderID))
		return pds, model.OrderStatusFailed
	} else {
		log.Info().Msg(fmt.Sprintf("order id : %d failed to refund", purchaseOrderID))
		return pds, model.OrderStatusProcessed
	}
}

// FindAll implements PurchaseOrderService
func (t *purchaseOrderServiceImpl) FindAll() []response.PurchaseOrderResponse {
	result, _ := t.PurchaseOrderRepository.FindAll()

	var PurchaseOrder []response.PurchaseOrderResponse
	for _, value := range result {
		PurchaseOrder = append(PurchaseOrder, value.ToResponse())
	}

	return PurchaseOrder
}

// FindByIdWithUser implements PurchaseOrderService
func (t *purchaseOrderServiceImpl) FindByIdWithUser(PurchaseOrderId uint32, userId uint32) (response.PurchaseOrderResponse, error) {
	purchaseOrderData, err := t.PurchaseOrderRepository.FindById(PurchaseOrderId)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}
	if purchaseOrderData.UserID != userId {
		return response.PurchaseOrderResponse{}, app_errors.NewMySQLNotFound(nil)
	}

	details, err := t.PurchaseOrderDetailService.FindByPurchaseOrderId(purchaseOrderData.ID)
	if err != nil {
		return response.PurchaseOrderResponse{}, err
	}
	poResponse := purchaseOrderData.ToResponse()
	poResponse.OrderDetails = details
	return poResponse, nil
}

func (t *purchaseOrderServiceImpl) MassRefund() error {
	poResults, err := t.PurchaseOrderRepository.GetAllPending()
	if err != nil {
		return err
	}
	for _, result := range poResults {
		_, err := t.RefundOrder(result.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
