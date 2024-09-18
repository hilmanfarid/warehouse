package repository

import (
	"errors"
	"golang-warehouse/data/request"
	"golang-warehouse/helper"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryRepository interface {
	Save(inventory model.Inventory) (model.Inventory, error)
	Update(inventory model.Inventory) (model.Inventory, error)
	Delete(inventoryId uint32)
	FindById(inventoryId uint32) (inventory model.Inventory, err error)
	FindAll() ([]model.Inventory, error)
	TransferStock(transferOrder model.TransferOrder) error
}

type InventoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewInventoryRepositoryImpl(Db *gorm.DB) InventoryRepository {
	return &InventoryRepositoryImpl{Db: Db}
}

// Delete implements InventoryRepository
func (t *InventoryRepositoryImpl) Delete(inventoryId uint32) {
	var inventory model.Inventory
	result := t.Db.Where("id = ?", inventoryId).Delete(&inventory)
	helper.ErrorPanic(result.Error)
}

// FindAll implements InventoryRepository
func (t *InventoryRepositoryImpl) FindAll() ([]model.Inventory, error) {
	var inventory []model.Inventory
	result := t.Db.Find(&inventory)
	if result != nil {
		return inventory, nil
	} else {
		return inventory, result.Error
	}
}

// FindById implements InventoryRepository
func (t *InventoryRepositoryImpl) FindById(inventoryId uint32) (inventory model.Inventory, err error) {
	_ = t.Db.Find(&inventory, inventoryId)
	if inventory != (model.Inventory{}) {
		return inventory, nil
	} else {
		return inventory, app_errors.NewMySQLNotFound(nil)
	}
}

// Save implements InventoryRepository
func (t *InventoryRepositoryImpl) Save(inventory model.Inventory) (model.Inventory, error) {
	result := t.Db.Create(&inventory)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return inventory, app_errors.NewMySQLDuplicateKey(result.Error)
	}
	if app_errors.IsMySQLReferenceError(result.Error) {
		return inventory, app_errors.NewGeneralError(result.Error)
	}

	return inventory, nil
}

// Update implements InventoryRepository
func (t *InventoryRepositoryImpl) Update(inventory model.Inventory) (model.Inventory, error) {
	result := t.Db.Model(&inventory).Updates(inventory)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return inventory, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return inventory, nil
}

func (t *InventoryRepositoryImpl) TransferStock(transferOrder model.TransferOrder) error {
	var sourceInventory model.Inventory
	var desinationInventory model.Inventory
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_id = ? and warehouse_id = ?", transferOrder.ProductID, transferOrder.SourceWarehouse).First(&sourceInventory)

		if result.Error != nil || sourceInventory.ID == 0 || sourceInventory.Quantity < transferOrder.Quantity {
			log.Error().Msg("invalid stock")
			return app_errors.NewGeneralError(errors.New("invalid stock"))
		}

		updateSourceParams := request.UpdateTransferOrderRequest{
			Quantity: sourceInventory.Quantity - transferOrder.Quantity,
		}
		tx.Model(&sourceInventory).Updates(updateSourceParams)

		result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_id = ? and warehouse_id = ?", transferOrder.ProductID, transferOrder.DestinationWarehouse).First(&desinationInventory)

		if result.Error != nil {
			log.Error().Msg("invalid destination inventory")
			return app_errors.NewGeneralError(errors.New("invalid stock"))
		}

		if desinationInventory.ID == 0 {
			_, err := t.Save(model.Inventory{
				ProductID:   transferOrder.ProductID,
				WarehouseID: transferOrder.DestinationWarehouse,
				Quantity:    transferOrder.Quantity,
				Status:      1,
			})
			if err != nil {
				return app_errors.NewGeneralError(err)
			}
		} else {
			updateDestParams := request.UpdateTransferOrderRequest{
				Quantity: desinationInventory.Quantity + transferOrder.Quantity,
			}
			tx.Model(&desinationInventory).Updates(updateDestParams)
		}
		return nil
	})
	return err
}
