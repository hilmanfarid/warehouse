package controller

import (
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/helper"
	"golang-warehouse/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PurchaseOrderController struct {
	purchaseOrderService service.PurchaseOrderService
}

func NewPurchaseOrderController(service service.PurchaseOrderService) *PurchaseOrderController {
	return &PurchaseOrderController{
		purchaseOrderService: service,
	}
}

func (controller *PurchaseOrderController) CreateOrder(ctx *gin.Context) {
	log.Info().Msg("create order")
	userInfo := parseUserClaim(ctx)
	createPurchaseOrderRequest := request.CreatePurchaseOrderRequest{
		UserID: userInfo.UserID,
		Status: 1,
	}
	err := ctx.ShouldBindJSON(&createPurchaseOrderRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.purchaseOrderService.CreateOrder(createPurchaseOrderRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	// since we are not using any background service to process the order,
	// assuming there is no issue after create, we will call process order directly

	result, err = controller.purchaseOrderService.ProcessOrder(result.ID)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PurchaseOrderController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid product")
	orderId := ctx.Param("id")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	userInfo := parseUserClaim(ctx)
	result, err := controller.purchaseOrderService.FindByIdWithUser(uint32(id), userInfo.UserID)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   result,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
