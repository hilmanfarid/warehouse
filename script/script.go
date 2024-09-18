package main

import (
	"fmt"
	"golang-warehouse/config"
	"golang-warehouse/repository"
	"golang-warehouse/service"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error loading .env file: %s", err))
	}
	db := config.DatabaseConnection()
	validate := validator.New()
	repo := repository.NewPurchaseOrderDetailRepositoryImpl(db)
	porepo := repository.NewPurchaseOrderRepositoryImpl(db)
	invrepo := repository.NewInventoryRepositoryImpl(db)
	prodrepo := repository.NewProductRepositoryImpl(db)
	productS := service.NewProductServiceImpl(prodrepo, validate)

	invService := service.NewInventoryServiceImpl(invrepo, validate)
	detailService := service.NewPurchaseOrderDetailServiceImpl(repo, porepo, invService, validate)
	poService := service.NewPurchaseOrderServiceImpl(productS, detailService, porepo, validate)

	detail, err := poService.RefundOrder(41)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(detail)
}
