package router

import (
	"golang-warehouse/controller"
	"golang-warehouse/controller/admin"
	"golang-warehouse/middleware"
	"golang-warehouse/service"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	authController *controller.AuthController,
	tokenService service.TokenService,
	productsController *admin.ProductController,
	shopController *admin.ShopController,
	warehouseController *admin.WarehouseController,
	inventoryController *admin.InventoryController,
	purchaserOrderController *controller.PurchaseOrderController,
	transferOrderController *admin.TransferOrderController,
) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")

	router.POST("/register", authController.Register)
	router.POST("/auth", authController.Login)

	productsRouter := baseRouter.Group("/admin/products")
	productsRouter.GET("", middleware.AuthUser(tokenService, "admin"), productsController.FindAll)
	productsRouter.GET("/:id", middleware.AuthUser(tokenService, "admin"), productsController.FindById)
	productsRouter.POST("", middleware.AuthUser(tokenService, "admin"), productsController.Create)
	productsRouter.PATCH("/:id", middleware.AuthUser(tokenService, "admin"), productsController.Update)
	productsRouter.DELETE("/:id", middleware.AuthUser(tokenService, "admin"), productsController.Delete)

	shopRouter := baseRouter.Group("/admin/shops")
	shopRouter.GET("", middleware.AuthUser(tokenService, "admin"), shopController.FindAll)
	shopRouter.GET("/:id", middleware.AuthUser(tokenService, "admin"), shopController.FindById)
	shopRouter.POST("", middleware.AuthUser(tokenService, "admin"), shopController.Create)
	shopRouter.PATCH("/:id", middleware.AuthUser(tokenService, "admin"), shopController.Update)
	shopRouter.DELETE("/:id", middleware.AuthUser(tokenService, "admin"), shopController.Delete)

	warehouseRouter := baseRouter.Group("/admin/warehouses")
	warehouseRouter.GET("", middleware.AuthUser(tokenService, "admin"), warehouseController.FindAll)
	warehouseRouter.GET("/:id", middleware.AuthUser(tokenService, "admin"), warehouseController.FindById)
	warehouseRouter.POST("", middleware.AuthUser(tokenService, "admin"), warehouseController.Create)
	warehouseRouter.PATCH("/:id", middleware.AuthUser(tokenService, "admin"), warehouseController.Update)
	warehouseRouter.DELETE("/:id", middleware.AuthUser(tokenService, "admin"), warehouseController.Delete)

	inventoryRouter := baseRouter.Group("/admin/inventory")
	inventoryRouter.GET("", middleware.AuthUser(tokenService, "admin"), inventoryController.FindAll)
	inventoryRouter.GET("/:id", middleware.AuthUser(tokenService, "admin"), inventoryController.FindById)
	inventoryRouter.POST("", middleware.AuthUser(tokenService, "admin"), inventoryController.Create)
	inventoryRouter.PATCH("/:id", middleware.AuthUser(tokenService, "admin"), inventoryController.Update)
	inventoryRouter.DELETE("/:id", middleware.AuthUser(tokenService, "admin"), inventoryController.Delete)

	purchaseOrderRouter := baseRouter.Group("/order")
	purchaseOrderRouter.POST("", middleware.AuthUser(tokenService, ""), purchaserOrderController.CreateOrder)
	purchaseOrderRouter.GET("/:id", middleware.AuthUser(tokenService, ""), purchaserOrderController.FindById)

	transferOrderRouter := baseRouter.Group("admin/transfer-order")
	transferOrderRouter.POST("", middleware.AuthUser(tokenService, ""), transferOrderController.TransferOrder)
	transferOrderRouter.GET("/:id", middleware.AuthUser(tokenService, ""), transferOrderController.FindById)

	return router
}
