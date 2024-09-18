package repository

import (
	"golang-warehouse/data/request"
	"golang-warehouse/helper"
	"golang-warehouse/model"
	app_errors "golang-warehouse/model/app_errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(userId uint32)
	FindById(userId uint32) (user model.User, err error)
	FindByEmail(username string) (user model.User, err error)
	FindAll() ([]model.User, error)
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (t *UserRepositoryImpl) FindByEmail(email string) (user model.User, err error) {
	_ = t.Db.Where("email = ?", email).First(&user)
	if user != (model.User{}) {
		return user, nil
	} else {
		return user, app_errors.NewMySQLNotFound(nil)
	}
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository
func (t *UserRepositoryImpl) Delete(userId uint32) {
	var user model.User
	result := t.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository
func (t *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var user []model.User
	result := t.Db.Find(&user)
	if result != nil {
		return user, nil
	} else {
		return user, result.Error
	}
}

// FindById implements UserRepository
func (t *UserRepositoryImpl) FindById(userId uint32) (user model.User, err error) {
	_ = t.Db.Find(&user, userId)
	if user != (model.User{}) {
		return user, nil
	} else {
		return user, app_errors.NewMySQLNotFound(nil)
	}
}

// Save implements UserRepository
func (t *UserRepositoryImpl) Save(user model.User) (model.User, error) {
	result := t.Db.Create(&user)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return user, app_errors.NewMySQLDuplicateKey(result.Error)
	} else if result.Error != nil {
		return user, app_errors.NewGeneralError(result.Error)
	}

	return user, nil
}

// Update implements UserRepository
func (t *UserRepositoryImpl) Update(user model.User) (model.User, error) {
	var updateTag = request.UpdateUserRequest{
		ID:    user.ID,
		Email: user.Email,
	}
	result := t.Db.Model(&user).Updates(updateTag)
	if app_errors.IsMySQLDuplicateKey(result.Error) {
		return user, app_errors.NewMySQLDuplicateKey(result.Error)
	}

	return user, nil
}
