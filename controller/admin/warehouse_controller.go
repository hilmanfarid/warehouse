package admin

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

type WarehouseController struct {
	warehouseService service.WarehouseService
}

func NewWarehouseController(service service.WarehouseService) *WarehouseController {
	return &WarehouseController{
		warehouseService: service,
	}
}

func (controller *WarehouseController) Create(ctx *gin.Context) {
	log.Info().Msg("create warehouse")
	createWarehouseRequest := request.CreateWarehouseRequest{Status: 1}
	err := ctx.ShouldBindJSON(&createWarehouseRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.warehouseService.Create(createWarehouseRequest)
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

func (controller *WarehouseController) Update(ctx *gin.Context) {
	log.Info().Msg("update warehouse")
	updateWarehouseRequest := request.UpdateWarehouseRequest{}
	err := ctx.ShouldBindJSON(&updateWarehouseRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	warehouseId := ctx.Param("id")
	id, err := strconv.Atoi(warehouseId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	updateWarehouseRequest.ID = uint32(id)

	result, err := controller.warehouseService.Update(updateWarehouseRequest)
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

func (controller *WarehouseController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete warehouse")
	warehouseId := ctx.Param("id")
	id, err := strconv.Atoi(warehouseId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	controller.warehouseService.Delete(uint32(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *WarehouseController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid warehouse")
	warehouseId := ctx.Param("id")
	id, err := strconv.Atoi(warehouseId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.warehouseService.FindById(uint32(id))
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

func (controller *WarehouseController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll warehouse")
	results := controller.warehouseService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   results,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
