package admin

import (
	"errors"
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/helper"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TransferOrderController struct {
	transferOrderService service.TransferOrderService
}

func NewTransferOrderController(service service.TransferOrderService) *TransferOrderController {
	return &TransferOrderController{
		transferOrderService: service,
	}
}

func (controller *TransferOrderController) TransferOrder(ctx *gin.Context) {
	log.Info().Msg("create order")
	userInfo := parseUserClaim(ctx)
	createTransferOrderRequest := request.CreateTransferOrderRequest{
		UserID: userInfo.UserID,
		Status: model.TransferOrderStatusCreated,
	}
	err := ctx.ShouldBindJSON(&createTransferOrderRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.transferOrderService.TransferOrder(createTransferOrderRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	// since we are not using any background service to process the order,
	// assuming there is no issue after create, we will call process order directly

	result, err = controller.transferOrderService.ProcessTransfer(result.ID)
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

func (controller *TransferOrderController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid product")
	tfId := ctx.Param("id")
	id, err := strconv.Atoi(tfId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.transferOrderService.FindById(uint32(id))
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

func parseUserClaim(ctx *gin.Context) *model.IDTokenCustomClaims {
	userClaim, ok := ctx.Get("userClaim")
	if !ok {
		helper.HTTPStatusError(ctx, app_errors.NewAuthorizationInvalid(errors.New("invalid user")))
		return nil
	}
	return userClaim.(*model.IDTokenCustomClaims)
}
