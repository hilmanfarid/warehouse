package repository

import (
	"golang-warehouse/data/request"
	"golang-warehouse/helper"
	"golang-warehouse/model"
	app_errors "golang-warehouse/model/app_errors"

	"gorm.io/gorm"
)

type ShopRepository interface {
	Save(shop model.Shop) (model.Shop, error)
	Update(shop model.Shop) (model.Shop, error)
	Delete(shopId uint32)
	FindById(shopId uint32) (shop model.Shop, err error)
	FindAll() ([]model.Shop, error)
}

type ShopRepositoryImpl struct {
	Db *gorm.DB
}

func NewShopRepositoryImpl(Db *gorm.DB) ShopRepository {
	return &ShopRepositoryImpl{Db: Db}
}

// Delete implements ShopRepository
func (t *ShopRepositoryImpl) Delete(shopId uint32) {
	var shop model.Shop
	result := t.Db.Where("id = ?", shopId).Delete(&shop)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ShopRepository
func (t *ShopRepositoryImpl) FindAll() ([]model.Shop, error) {
	var shop []model.Shop
	result := t.Db.Find(&shop)
	if result != nil {
		return shop, nil
	} else {
		return shop, result.Error
	}
}

// FindById implements ShopRepository
func (t *ShopRepositoryImpl) FindById(shopId uint32) (shop model.Shop, err error) {
	_ = t.Db.Find(&shop, shopId)
	if shop != (model.Shop{}) {
		return shop, nil
	} else {
		return shop, app_errors.NewMySQLNotFound(nil)
	}
}

// Save implements ShopRepository
func (t *ShopRepositoryImpl) Save(shop model.Shop) (model.Shop, error) {
	result := t.Db.Create(&shop)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return shop, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return shop, nil
}

// Update implements ShopRepository
func (t *ShopRepositoryImpl) Update(shop model.Shop) (model.Shop, error) {
	var updateTag = request.UpdateShopRequest{
		ID:   shop.ID,
		Name: shop.Name,
	}
	result := t.Db.Model(&shop).Updates(updateTag)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return shop, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return shop, nil
}
