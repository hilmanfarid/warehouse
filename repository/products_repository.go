package repository

import (
	"golang-warehouse/helper"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(productId uint32)
	FindById(productId uint32) (product model.Product, err error)
	FindAll() ([]model.Product, error)
}

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

// Delete implements ProductRepository
func (t *ProductRepositoryImpl) Delete(productId uint32) {
	var product model.Product
	result := t.Db.Where("id = ?", productId).Delete(&product)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ProductRepository
func (t *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	var product []model.Product
	result := t.Db.Find(&product)
	if result != nil {
		return product, nil
	} else {
		return product, result.Error
	}
}

// FindById implements ProductRepository
func (t *ProductRepositoryImpl) FindById(productId uint32) (product model.Product, err error) {
	_ = t.Db.Find(&product, productId)
	if product != (model.Product{}) {
		return product, nil
	} else {
		return product, app_errors.NewMySQLNotFound(nil)
	}
}

// Save implements ProductRepository
func (t *ProductRepositoryImpl) Save(product model.Product) (model.Product, error) {
	result := t.Db.Create(&product)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return product, app_errors.NewMySQLDuplicateKey(result.Error)
	}
	if app_errors.IsMySQLReferenceError(result.Error) {
		return product, app_errors.NewGeneralError(result.Error)
	}

	return product, nil
}

// Update implements ProductRepository
func (t *ProductRepositoryImpl) Update(product model.Product) (model.Product, error) {
	result := t.Db.Model(&product).Updates(product)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return product, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return product, nil
}
