package repository

import (
	"golang-warehouse/helper"
	"golang-warehouse/model"
	app_errors "golang-warehouse/model/app_errors"

	"gorm.io/gorm"
)

type WarehouseRepository interface {
	Save(warehouse model.Warehouse) (model.Warehouse, error)
	Update(warehouse model.Warehouse) (model.Warehouse, error)
	Delete(warehouseId uint32)
	FindById(warehouseId uint32) (warehouse model.Warehouse, err error)
	FindAll() ([]model.Warehouse, error)
}

type WarehouseRepositoryImpl struct {
	Db *gorm.DB
}

func NewWarehouseRepositoryImpl(Db *gorm.DB) WarehouseRepository {
	return &WarehouseRepositoryImpl{Db: Db}
}

// Delete implements WarehouseRepository
func (t *WarehouseRepositoryImpl) Delete(warehouseId uint32) {
	var warehouse model.Warehouse
	result := t.Db.Where("id = ?", warehouseId).Delete(&warehouse)
	helper.ErrorPanic(result.Error)
}

// FindAll implements WarehouseRepository
func (t *WarehouseRepositoryImpl) FindAll() ([]model.Warehouse, error) {
	var warehouse []model.Warehouse
	result := t.Db.Find(&warehouse)
	if result != nil {
		return warehouse, nil
	} else {
		return warehouse, result.Error
	}
}

// FindById implements WarehouseRepository
func (t *WarehouseRepositoryImpl) FindById(warehouseId uint32) (warehouse model.Warehouse, err error) {
	_ = t.Db.Find(&warehouse, warehouseId)
	if warehouse != (model.Warehouse{}) {
		return warehouse, nil
	} else {
		return warehouse, app_errors.NewMySQLNotFound(nil)
	}
}

// Save implements WarehouseRepository
func (t *WarehouseRepositoryImpl) Save(warehouse model.Warehouse) (model.Warehouse, error) {
	result := t.Db.Create(&warehouse)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return warehouse, app_errors.NewMySQLDuplicateKey(result.Error)
	}
	if app_errors.IsMySQLReferenceError(result.Error) {
		return warehouse, app_errors.NewGeneralError(result.Error)
	}

	return warehouse, nil
}

// Update implements WarehouseRepository
func (t *WarehouseRepositoryImpl) Update(warehouse model.Warehouse) (model.Warehouse, error) {
	result := t.Db.Model(&warehouse).Updates(warehouse)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return warehouse, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return warehouse, nil
}
