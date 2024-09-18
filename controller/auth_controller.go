package controller

import (
	"errors"
	"golang-warehouse/data/request"
	"golang-warehouse/data/response"
	"golang-warehouse/helper"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AuthController struct {
	UserService  service.UserService
	TokenService service.TokenService
}

func NewAuthController(service service.UserService, tokenService service.TokenService) *AuthController {
	return &AuthController{
		UserService:  service,
		TokenService: tokenService,
	}
}

func (controller *AuthController) Register(ctx *gin.Context) {
	log.Info().Msg("register user")
	createAuthRequest := request.CreateUserRequest{Status: 1}
	err := ctx.ShouldBindJSON(&createAuthRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	err = controller.UserService.Signup(createAuthRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthController) Login(ctx *gin.Context) {
	log.Info().Msg("login user")
	loginRequest := request.LoginUserRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}

	user, err := controller.UserService.Signin(loginRequest)
	if err != nil {
		helper.HTTPStatusError(ctx, err)
		return
	}
	_, err = controller.TokenService.NewPairFromUser(&user)
	if err != nil {
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: user.Token,
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
