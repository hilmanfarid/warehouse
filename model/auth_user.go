package model

import "github.com/golang-jwt/jwt/v5"

type IDTokenCustomClaims struct {
	UserID uint32 `json:"user_id"`
	Email  string `json:"email"`
	Scope  string `json:"scope"`

	jwt.RegisteredClaims
}
