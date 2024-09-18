package repository

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/helper"
	"golang-warehouse/model"
	app_errors "golang-warehouse/model/app_errors"

	"gorm.io/gorm"
)

type PurchaseOrderRepository interface {
	CreateOrder(purchaseOrder model.PurchaseOrder, details []model.PurchaseOrderDetail) (response.PurchaseOrderResponse, error)
	Save(purchaseOrder model.PurchaseOrder) (model.PurchaseOrder, error)
	Update(purchaseOrder model.PurchaseOrder) (model.PurchaseOrder, error)
	Delete(purchaseOrderId uint32)
	FindById(purchaseOrderId uint32) (purchaseOrder model.PurchaseOrder, err error)
	FindAll() ([]model.PurchaseOrder, error)
	GetAllPending() ([]model.PurchaseOrder, error)
}

type PurchaseOrderRepositoryImpl struct {
	Db *gorm.DB
}

func NewPurchaseOrderRepositoryImpl(Db *gorm.DB) PurchaseOrderRepository {
	return &PurchaseOrderRepositoryImpl{Db: Db}
}

// Delete implements PurchaseOrderRepository
func (t *PurchaseOrderRepositoryImpl) Delete(purchaseOrderId uint32) {
	var purchaseOrder model.PurchaseOrder
	result := t.Db.Where("id = ?", purchaseOrderId).Delete(&purchaseOrder)
	helper.ErrorPanic(result.Error)
}

// FindAll implements PurchaseOrderRepository
func (t *PurchaseOrderRepositoryImpl) FindAll() ([]model.PurchaseOrder, error) {
	var purchaseOrder []model.PurchaseOrder
	result := t.Db.Find(&purchaseOrder)
	if result != nil {
		return purchaseOrder, nil
	} else {
		return purchaseOrder, result.Error
	}
}

// FindById implements PurchaseOrderRepository
func (t *PurchaseOrderRepositoryImpl) FindById(purchaseOrderId uint32) (purchaseOrder model.PurchaseOrder, err error) {
	result := t.Db.Find(&purchaseOrder, purchaseOrderId)
	if result.Error != nil {
		return purchaseOrder, app_errors.NewGeneralError(result.Error)
	} else {
		return purchaseOrder, nil
	}
}

// CreateOrder save order along with details
func (t *PurchaseOrderRepositoryImpl) CreateOrder(purchaseOrder model.PurchaseOrder, details []model.PurchaseOrderDetail) (response.PurchaseOrderResponse, error) {
	var purchaseResponse response.PurchaseOrderResponse
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		result := t.Db.Create(&purchaseOrder)
		if result.Error != nil {
			return app_errors.NewGeneralError(result.Error)
		}
		for _, detail := range details {
			detail.PurchaseOrderID = purchaseOrder.ID
			detail.Status = model.OrderDetailStatusCreated
			result = t.Db.Create(&detail)
			if app_errors.IsMySQLDuplicateKey(result.Error) {
				return app_errors.NewMySQLDuplicateKey(result.Error)
			}
		}
		return nil
	})
	purchaseResponse = purchaseOrder.ToResponse()
	for _, detail := range details {
		purchaseResponse.OrderDetails = append(purchaseResponse.OrderDetails, detail.ToResponse())
	}
	return purchaseResponse, err
}

// Save implements PurchaseOrderRepository
func (t *PurchaseOrderRepositoryImpl) Save(purchaseOrder model.PurchaseOrder) (model.PurchaseOrder, error) {
	result := t.Db.Create(&purchaseOrder)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return purchaseOrder, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return purchaseOrder, nil
}

// Update implements PurchaseOrderRepository
func (t *PurchaseOrderRepositoryImpl) Update(purchaseOrder model.PurchaseOrder) (model.PurchaseOrder, error) {
	var updateTag = request.UpdatePurchaseOrderRequest{
		ID:          purchaseOrder.ID,
		UserID:      purchaseOrder.UserID,
		ShopID:      purchaseOrder.ShopID,
		Status:      purchaseOrder.Status,
		TotalAmount: purchaseOrder.TotalAmount,
		ProcessAt:   purchaseOrder.ProcessAt,
		SuccessAt:   purchaseOrder.SuccessAt,
		FailedAt:    purchaseOrder.FailedAt,
	}
	result := t.Db.Model(&purchaseOrder).Updates(updateTag)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return purchaseOrder, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return purchaseOrder, nil
}

func (t *PurchaseOrderRepositoryImpl) GetAllPending() ([]model.PurchaseOrder, error) {
	var purchaseOrder []model.PurchaseOrder

	// get pending after 1 minute up to 1 day
	result := t.Db.Where("status = ?", model.OrderStatusProcessed).
		Where("TIMESTAMPDIFF(MINUTE,created_at,NOW()) < 1440 and TIMESTAMPDIFF(MINUTE,created_at,NOW()) >= 1").
		Limit(100).
		Find(&purchaseOrder)
	if result != nil && result.Error == nil {
		return purchaseOrder, nil
	} else {
		return purchaseOrder, app_errors.NewGeneralError(result.Error)
	}
}
