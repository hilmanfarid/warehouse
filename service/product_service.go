package service

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	Create(product request.CreateProductRequest) (response.ProductResponse, error)
	Update(product request.UpdateProductRequest) (response.ProductResponse, error)
	Delete(product uint32)
	FindById(product uint32) (response.ProductResponse, error)
	FindAll() []response.ProductResponse
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &productServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

// Create implements ProductService
func (t *productServiceImpl) Create(product request.CreateProductRequest) (response.ProductResponse, error) {
	err := t.Validate.Struct(product)
	if err != nil {
		return response.ProductResponse{}, err
	}
	params := model.Product{
		Name:   product.Name,
		Code:   product.Code,
		Status: product.Status,
		Price:  product.Price,
	}
	result, err := t.ProductRepository.Save(params)
	if err != nil {
		return response.ProductResponse{}, err
	}
	return result.ToResponse(), nil
}

// Delete implements ProductService
func (t *productServiceImpl) Delete(productId uint32) {
	t.ProductRepository.Delete(productId)
}

// FindAll implements ProductService
func (t *productServiceImpl) FindAll() []response.ProductResponse {
	result, _ := t.ProductRepository.FindAll()

	var products []response.ProductResponse
	for _, value := range result {
		products = append(products, value.ToResponse())
	}
	return products
}

// FindById implements ProductService
func (t *productServiceImpl) FindById(productId uint32) (response.ProductResponse, error) {
	productData, err := t.ProductRepository.FindById(productId)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return productData.ToResponse(), nil
}

// Update implements ProductService
func (t *productServiceImpl) Update(product request.UpdateProductRequest) (response.ProductResponse, error) {
	updatedObj, err := t.ProductRepository.FindById(product.ID)
	if err != nil {
		return response.ProductResponse{}, app_errors.NewMySQLNotFound(err)
	}

	if product.Name != "" {
		updatedObj.Name = product.Name
	}
	if product.Status != 0 {
		updatedObj.Status = product.Status
	}
	if product.Code != "" {
		updatedObj.Code = product.Code
	}
	if product.Price != 0 {
		updatedObj.Price = product.Price
	}

	result, err := t.ProductRepository.Update(updatedObj)
	if err != nil {
		return response.ProductResponse{}, nil
	}

	return result.ToResponse(), nil
}
