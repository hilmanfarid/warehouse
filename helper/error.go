package helper

import (
	"golang-warehouse/data/response"
	"golang-warehouse/model/app_errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func HTTPStatusError(ctx *gin.Context, err error) {
	var baseError *app_errors.BaseError
	baseError, ok := err.(*app_errors.BaseError)
	if !ok {
		baseError = app_errors.NewGeneralError(err)
	}
	code, _ := strconv.Atoi(baseError.Detail().Code)
	webResponse := response.Response{
		Code:    code,
		Status:  "Failed",
		Message: baseError.Detail().Message,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, webResponse)
	logger := log.With().Str("stack_trace", baseError.StackTrace()).Logger()
	logger.Error().Msg(baseError.Error())
}
