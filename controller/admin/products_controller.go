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

type ProductController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (controller *ProductController) Create(ctx *gin.Context) {
	log.Info().Msg("create product")
	createProductRequest := request.CreateProductRequest{Status: 1}
	err := ctx.ShouldBindJSON(&createProductRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.productService.Create(createProductRequest)
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

func (controller *ProductController) Update(ctx *gin.Context) {
	log.Info().Msg("update product")
	updateProductRequest := request.UpdateProductRequest{}
	err := ctx.ShouldBindJSON(&updateProductRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	updateProductRequest.ID = uint32(id)

	result, err := controller.productService.Update(updateProductRequest)
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

func (controller *ProductController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete product")
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	controller.productService.Delete(uint32(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid product")
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	result, err := controller.productService.FindById(uint32(id))
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

func (controller *ProductController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll product")
	results := controller.productService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   results,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
