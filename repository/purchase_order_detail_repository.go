package repository

import (
	"errors"
	"golang-warehouse/data/request"
	"golang-warehouse/model"
	app_errors "golang-warehouse/model/app_errors"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PurchaseOrderDetailRepository interface {
	Save(purchaseOrderDetail model.PurchaseOrderDetail) (model.PurchaseOrderDetail, error)
	PurchaseStock(purchaseOrder model.PurchaseOrder, purchaseOrderDetail model.PurchaseOrderDetail, state model.OrderDetailStatus) (model.PurchaseOrderDetail, error)
	Delete(purchaseOrderDetailId uint32) *app_errors.BaseError
	FindById(purchaseOrderDetailId uint32) (purchaseOrderDetail model.PurchaseOrderDetail, err error)
	FindAll() ([]model.PurchaseOrderDetail, error)
	FindByPurchaseOrderID(purchaseOrderID uint32) ([]model.PurchaseOrderDetail, error)
}

type PurchaseOrderDetailRepositoryImpl struct {
	Db *gorm.DB
}

func NewPurchaseOrderDetailRepositoryImpl(Db *gorm.DB) PurchaseOrderDetailRepository {
	return &PurchaseOrderDetailRepositoryImpl{Db: Db}
}

// Delete implements PurchaseOrderDetailRepository
func (t *PurchaseOrderDetailRepositoryImpl) Delete(purchaseOrderDetailId uint32) *app_errors.BaseError {
	var purchaseOrderDetail model.PurchaseOrderDetail
	result := t.Db.Where("id = ?", purchaseOrderDetailId).Delete(&purchaseOrderDetail)
	if result.Error != nil {
		return app_errors.NewGeneralError(result.Error)
	}
	return nil
}

// FindAll implements PurchaseOrderDetailRepository
func (t *PurchaseOrderDetailRepositoryImpl) FindAll() ([]model.PurchaseOrderDetail, error) {
	var purchaseOrderDetail []model.PurchaseOrderDetail
	result := t.Db.Find(&purchaseOrderDetail)
	if result != nil {
		return purchaseOrderDetail, nil
	} else {
		return purchaseOrderDetail, result.Error
	}
}

// FindById implements PurchaseOrderDetailRepository
func (t *PurchaseOrderDetailRepositoryImpl) FindById(purchaseOrderDetailId uint32) (purchaseOrderDetail model.PurchaseOrderDetail, err error) {
	_ = t.Db.Find(&purchaseOrderDetail, purchaseOrderDetailId)
	if purchaseOrderDetail != (model.PurchaseOrderDetail{}) {
		return purchaseOrderDetail, nil
	} else {
		return purchaseOrderDetail, app_errors.NewMySQLNotFound(nil)
	}
}

func (t *PurchaseOrderDetailRepositoryImpl) FindByPurchaseOrderID(purchaseOrderID uint32) ([]model.PurchaseOrderDetail, error) {
	var purchaseOrderDetail []model.PurchaseOrderDetail
	result := t.Db.Where("purchase_order_id = ?", purchaseOrderID).Find(&purchaseOrderDetail)
	if result.Error != nil {
		return purchaseOrderDetail, app_errors.NewGeneralError(result.Error)
	}
	return purchaseOrderDetail, nil
}

// Save implements PurchaseOrderDetailRepository
func (t *PurchaseOrderDetailRepositoryImpl) Save(purchaseOrderDetail model.PurchaseOrderDetail) (model.PurchaseOrderDetail, error) {
	result := t.Db.Create(&purchaseOrderDetail)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return purchaseOrderDetail, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return purchaseOrderDetail, nil
}

// PurchaseStock get stock from inventory
func (t *PurchaseOrderDetailRepositoryImpl) PurchaseStock(purchaseOrder model.PurchaseOrder, purchaseOrderDetail model.PurchaseOrderDetail, state model.OrderDetailStatus) (model.PurchaseOrderDetail, error) {
	var inventoryResult model.Inventory
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		query := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Joins("left join warehouses on inventory.warehouse_id = warehouses.id").
			Where("product_id = ? and warehouses.shop_id = ?", purchaseOrderDetail.ProductID, purchaseOrder.ShopID).
			Where("inventory.status = ?", model.StatusAvailable)
		var result *gorm.DB
		if state == model.OrderDetailStatusSucceeded {
			query = query.Where("quantity > ?", purchaseOrderDetail.Quantity)
		} else if state == model.OrderDetailStatusRefunded {
			query = query.Where("inventory.warehouse_id", purchaseOrderDetail.WarehouseID)
		}

		query.First(&inventoryResult)
		if inventoryResult.ID == 0 {
			log.Error().Msg("invalid stock")
			return app_errors.NewGeneralError(errors.New("invalid stock"))
		}

		if state == model.OrderDetailStatusSucceeded {
			inventoryResult.Quantity -= purchaseOrderDetail.Quantity // success
		} else {
			state = model.OrderDetailStatusRefunded
			inventoryResult.Quantity += purchaseOrderDetail.Quantity // failed and refund stock
		}
		result = tx.Model(&inventoryResult).Updates(inventoryResult)
		if result.Error != nil {
			return app_errors.NewGeneralError(result.Error)
		}

		var updateParam = request.UpdatePurchaseOrderDetailsRequest{
			Status:    state,
			SuccessAt: time.Now(),
		}
		if state == model.OrderDetailStatusSucceeded {
			updateParam.SuccessAt = time.Now()
			updateParam.WarehouseID = inventoryResult.WarehouseID
		} else {
			updateParam.FailedAt = time.Now()
		}

		result = tx.Model(&purchaseOrderDetail).Updates(updateParam)
		if result.Error != nil {
			return app_errors.NewGeneralError(result.Error)
		}
		return nil
	})
	if err != nil {
		var updateParam = request.UpdatePurchaseOrderDetailsRequest{
			Status:   model.OrderDetailStatusFailed,
			FailedAt: time.Now(),
		}
		result := t.Db.Model(&purchaseOrderDetail).Updates(updateParam)
		if result.Error != nil {
			return purchaseOrderDetail, app_errors.NewGeneralError(result.Error)
		}
	}
	return purchaseOrderDetail, nil
}

//
//func (t *PurchaseOrderDetailRepositoryImpl) GetAllWarehouseID(shopID uint32) ([]model.Warehouse, error) {
//	var warehouseData []int
//	//var warehouse model.Warehouse
//	result := t.Db.Select("id").Where("shop_id = ?", shopID).Find(&warehouseData)
//	if result.Error != nil {
//		fmt.Println("ererrrrrrrrrrrrrrrrrrrrrrrr")
//		return nil, app_errors.NewGeneralError(result.Error)
//	}
//
//	return warehouseData, nil
//}
