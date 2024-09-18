package repository

import (
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"

	"gorm.io/gorm"
)

type TransferOrderRepository interface {
	Save(transferOrder model.TransferOrder) (model.TransferOrder, error)
	FindById(transferOrderId uint32) (transferOrder model.TransferOrder, err error)
	Update(transferOrder model.TransferOrder) (model.TransferOrder, error)
}

type TransferOrderRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransferOrderRepositoryImpl(Db *gorm.DB) TransferOrderRepository {
	return &TransferOrderRepositoryImpl{Db: Db}
}

// FindById implements TransferOrderRepository
func (t *TransferOrderRepositoryImpl) FindById(transferOrderId uint32) (transferOrder model.TransferOrder, err error) {
	_ = t.Db.Find(&transferOrder, transferOrderId)
	if transferOrder != (model.TransferOrder{}) {
		return transferOrder, nil
	} else {
		return transferOrder, app_errors.NewMySQLNotFound(nil)
	}
}

// Save implements TransferOrderRepository
func (t *TransferOrderRepositoryImpl) Save(transferOrder model.TransferOrder) (model.TransferOrder, error) {
	result := t.Db.Create(&transferOrder)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return transferOrder, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return transferOrder, nil
}

func (t *TransferOrderRepositoryImpl) Update(transferOrder model.TransferOrder) (model.TransferOrder, error) {
	result := t.Db.Model(&transferOrder).Updates(transferOrder)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return transferOrder, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return transferOrder, nil
}
