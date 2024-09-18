package main

import (
	"fmt"
	"golang-warehouse/config"
	"golang-warehouse/controller"
	"golang-warehouse/controller/admin"
	_ "golang-warehouse/docs"
	"golang-warehouse/helper"
	"golang-warehouse/repository"
	"golang-warehouse/router"
	"golang-warehouse/service"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {
	log.Info().Msg("Started Server!")
	if os.Getenv("ENVIRONMENT") != "dev" && os.Getenv("ENVIRONMENT") != "stg" && os.Getenv("ENVIRONMENT") != "prod" {
		_ = godotenv.Load(".env")
	}

	time.Local = time.FixedZone("Asia/Jakarta", int((7 * time.Hour).Seconds()))
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	// Repository
	productRepository := repository.NewProductRepositoryImpl(db)
	shopRespository := repository.NewShopRepositoryImpl(db)
	warehouseRespository := repository.NewWarehouseRepositoryImpl(db)
	inventoryRespository := repository.NewInventoryRepositoryImpl(db)
	userRepository := repository.NewUserRepositoryImpl(db)
	purchaseOrderRepository := repository.NewPurchaseOrderRepositoryImpl(db)
	purchaseOrderDetailRepository := repository.NewPurchaseOrderDetailRepositoryImpl(db)
	transferOrderRepository := repository.NewTransferOrderRepositoryImpl(db)

	// Service
	productService := service.NewProductServiceImpl(productRepository, validate)
	shopService := service.NewShopServiceImpl(shopRespository, validate)
	warehouseService := service.NewWarehouseServiceImpl(warehouseRespository, validate)
	inventoryService := service.NewInventoryServiceImpl(inventoryRespository, validate)
	tokenService := service.NewTokenServiceImpl(userRepository, 86400)
	userService := service.NewUserService(userRepository, validate)
	purchaseOrderdetailService := service.NewPurchaseOrderDetailServiceImpl(purchaseOrderDetailRepository, purchaseOrderRepository, inventoryService, validate)
	purchaseOrderService := service.NewPurchaseOrderServiceImpl(productService, purchaseOrderdetailService, purchaseOrderRepository, validate)
	transferOrderService := service.NewTransferOrderServiceImpl(inventoryRespository, transferOrderRepository, validate)

	// Controller
	authController := controller.NewAuthController(userService, tokenService)
	productsController := admin.NewProductController(productService)
	shopController := admin.NewShopController(shopService)
	warehouseController := admin.NewWarehouseController(warehouseService)
	inventoryController := admin.NewInventoryController(inventoryService)
	purchaserOrderController := controller.NewPurchaseOrderController(purchaseOrderService)
	transferOrderController := admin.NewTransferOrderController(transferOrderService)

	// Router
	routes := router.NewRouter(
		authController,
		tokenService,
		productsController,
		shopController,
		warehouseController,
		inventoryController,
		purchaserOrderController,
		transferOrderController,
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
