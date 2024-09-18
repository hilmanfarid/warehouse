package service

import (
	"fmt"
	"golang-warehouse/model"
	"golang-warehouse/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	NewPairFromUser(u *model.User) (*jwt.Token, error)
	ValidateIDToken(tokenString string) (*model.IDTokenCustomClaims, error)
}

type tokenServiceImpl struct {
	UserRepository repository.UserRepository
	ExpirationSecs int64
}

func (s tokenServiceImpl) NewPairFromUser(u *model.User) (*jwt.Token, error) {
	var (
		t            *jwt.Token
		signedString string
	)
	user := model.IDTokenCustomClaims{
		UserID: u.ID,
		Email:  u.Email,
		Scope:  u.Role,
	}
	expires := jwt.NumericDate{Time: time.Now().Add(time.Duration(s.ExpirationSecs * 1000000000))}
	user.ExpiresAt = &expires
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	signedString, err := t.SignedString([]byte(u.Secret))
	if err != nil {
		return nil, err
	}

	u.Token = signedString
	return t, nil
}

func NewTokenServiceImpl(userRepository repository.UserRepository, ex int) TokenService {
	return &tokenServiceImpl{
		UserRepository: userRepository,
		ExpirationSecs: int64(ex),
	}
}

// ValidateIDToken validates the id token jwt string
// It returns the user extract from the IDTokenCustomClaims
func (s tokenServiceImpl) ValidateIDToken(tokenString string) (*model.IDTokenCustomClaims, error) {
	claims := &model.IDTokenCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		user, err := s.UserRepository.FindByEmail(claims.Email)
		if err != nil {
			return nil, err
		}
		fmt.Println([]byte(user.Secret))
		return []byte(user.Secret), nil
	})

	// For now we'll just return the error and handle logging in service level
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("ID token is invalid")
	}

	claims, ok := token.Claims.(*model.IDTokenCustomClaims)
	if !ok {
		return nil, fmt.Errorf("ID token valid but couldn't parse claims")
	}

	if claims.ExpiresAt == nil {
		return nil, fmt.Errorf("ID token is invalid")
	}

	if claims.ExpiresAt.Sub(time.Now()) < 0 {
		return nil, fmt.Errorf("token Expired")
	}
	return claims, nil
}
