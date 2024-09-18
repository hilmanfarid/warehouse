package service

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang-warehouse/data/request"
	"golang-warehouse/model"
	"golang-warehouse/model/app_errors"
	"golang-warehouse/repository"
	"net/http"
	"net/mail"
	"strconv"

	passValidator "github.com/go-passwd/validator"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Get(id uint32) (model.User, error)
	Signup(u request.CreateUserRequest) error
	Signin(userRequest request.LoginUserRequest) (model.User, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Get retrieves a user based on their uuid
func (s *userServiceImpl) Get(id uint32) (model.User, error) {
	u, err := s.UserRepository.FindById(id)

	return u, err
}

// Signup reaches our to a UserRepository to verify the
// email address is available and signs up the user if this is the case
func (s *userServiceImpl) Signup(u request.CreateUserRequest) error {
	secret, _ := randomHex(16)
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return app_errors.NewInvalidEmailError(err)
	}
	passwordValidator := passValidator.CommonPassword(nil)
	err = passwordValidator(u.Password)

	if err != nil {
		detail := &app_errors.ErrorDetails{
			OrinalError: err.Error(),
			Message:     "password too simple",
			Code:        strconv.Itoa(http.StatusUnauthorized),
		}
		return app_errors.NewBaseError(detail)
	}

	password := encryptSha256([]byte(u.Password), []byte(secret))

	// now I realize why I originally used Signup(ctx, email, password)
	// then created a user. It's somewhat un-natural to mutate the user here
	user := model.User{
		Email:           u.Email,
		CryptedPassword: password,
		Secret:          secret,
		Token:           "",
		Status:          2,
		Role:            "user",
	}

	if _, err := s.UserRepository.Save(user); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) Signin(userRequest request.LoginUserRequest) (model.User, error) {
	uFetched, err := s.UserRepository.FindByEmail(userRequest.Email)
	if err != nil {
		return model.User{}, app_errors.AuthorizationInvalidError
	}

	fmt.Println(uFetched.CryptedPassword)
	messageMAC, _ := hex.DecodeString(uFetched.CryptedPassword)

	match := ValidMAC([]byte(userRequest.Password), messageMAC, []byte(uFetched.Secret))
	if !match {
		return model.User{}, app_errors.AuthorizationInvalidError
	}

	return uFetched, nil
}

func encryptSha256(message, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	encrpytedString := mac.Sum(nil)
	return hex.EncodeToString(encrpytedString)
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
