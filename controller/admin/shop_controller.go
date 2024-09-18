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

type ShopController struct {
	shopService service.ShopService
}

func NewShopController(service service.ShopService) *ShopController {
	return &ShopController{
		shopService: service,
	}
}

func (controller *ShopController) Create(ctx *gin.Context) {
	log.Info().Msg("create shop")
	createShopRequest := request.CreateShopRequest{Status: 1}
	err := ctx.ShouldBindJSON(&createShopRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.shopService.Create(createShopRequest)
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

func (controller *ShopController) Update(ctx *gin.Context) {
	log.Info().Msg("update shop")
	updateShopRequest := request.UpdateShopRequest{}
	err := ctx.ShouldBindJSON(&updateShopRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	shopId := ctx.Param("id")
	id, err := strconv.Atoi(shopId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	updateShopRequest.ID = uint32(id)

	result, err := controller.shopService.Update(updateShopRequest)
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

func (controller *ShopController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete shop")
	shopId := ctx.Param("id")
	id, err := strconv.Atoi(shopId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	controller.shopService.Delete(uint32(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ShopController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid shop")
	shopId := ctx.Param("id")
	id, err := strconv.Atoi(shopId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.shopService.FindById(uint32(id))
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

func (controller *ShopController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll shop")
	results := controller.shopService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   results,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
