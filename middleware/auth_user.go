package middleware

import (
	apperror "golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// IMPORTS OMITTED - Make sure to import validator/v10
// My auto import always uses V9

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// AuthUser extracts a user from the Authorization header
// which is of the form "Bearer token"
// It sets the user to the context if the user exists
func AuthUser(s service.TokenService, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		// bind Authorization Header to h and check for validation errors
		if err := c.ShouldBindHeader(&h); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				// we used this type in bind_data to extract desired fields from errs
				// you might consider extracting it
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				err := apperror.NewBadRequest("Invalid request parameters. See invalidArgs")

				c.JSON(err.Status(), gin.H{
					"error":       err,
					"invalidArgs": invalidArgs,
				})
				c.Abort()
				return
			}

			// otherwise error type is unknown
			err := apperror.NewInternal()
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			err := apperror.NewAuthorization("Must provide Authorization header with format `Bearer {token}`")

			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		// validate ID token here
		userClaim, err := s.ValidateIDToken(idTokenHeader[1])
		if err != nil {
			baseError := app_errors.NewGeneralError(err)
			logger := log.With().Str("stack_trace", baseError.StackTrace()).Logger()
			logger.Error().Msg(baseError.Error())

			err := apperror.NewAuthorization("Provided token is invalid")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		if role == "admin" && role != userClaim.Scope {
			err := apperror.NewBadRequest("Forbidden")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		c.Set("userClaim", userClaim)
		c.Next()
	}
}
