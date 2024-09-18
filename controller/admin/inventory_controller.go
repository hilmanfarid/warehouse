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

type InventoryController struct {
	inventoryService service.InventoryService
}

func NewInventoryController(service service.InventoryService) *InventoryController {
	return &InventoryController{
		inventoryService: service,
	}
}

func (controller *InventoryController) Create(ctx *gin.Context) {
	log.Info().Msg("create inventory")
	createInventoryRequest := request.CreateInventoryRequest{Status: 1}
	err := ctx.ShouldBindJSON(&createInventoryRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.inventoryService.Create(createInventoryRequest)
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

func (controller *InventoryController) Update(ctx *gin.Context) {
	log.Info().Msg("update inventory")
	updateInventoryRequest := request.UpdateInventoryRequest{}
	err := ctx.ShouldBindJSON(&updateInventoryRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	inventoryId := ctx.Param("id")
	id, err := strconv.Atoi(inventoryId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	updateInventoryRequest.ID = uint32(id)

	result, err := controller.inventoryService.Update(updateInventoryRequest)
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

func (controller *InventoryController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete inventory")
	inventoryId := ctx.Param("id")
	id, err := strconv.Atoi(inventoryId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	controller.inventoryService.Delete(uint32(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *InventoryController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid inventory")
	inventoryId := ctx.Param("id")
	id, err := strconv.Atoi(inventoryId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.inventoryService.FindById(uint32(id))
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

func (controller *InventoryController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll inventory")
	results := controller.inventoryService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   results,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
