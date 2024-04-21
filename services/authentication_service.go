package services

import (
	"context"
	"errors"
	"github.com/SawitProRecruitment/UserService/forms"
	"github.com/SawitProRecruitment/UserService/modules"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthenticationService struct {
	repository   repository.UserRepositoryInterface
	passwordAuth modules.PasswordAuthInterface
	jwtAuth      modules.JsonWebTokenUtilInterface
}

func (a AuthenticationService) Authenticate(form forms.UserLoginForm) (*AuthenticationResult, error) {

	ctx := context.Background()

	result := &AuthenticationResult{
		ValidationErrors: nil,
		IsSuccess:        false,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(form)

	if err != nil {

		var validationErrors validator.ValidationErrors

		errors.As(err, &validationErrors)

		validationErrorMessages := utils.CollectValidationErrorMessages(form, validationErrors)

		result.HasValidationErrors = true
		result.ValidationErrors = validationErrorMessages

		return result, nil
	}

	getByPhoneNumberInput := repository.GetUserByPhoneNumberInput{
		PhoneNumber: form.PhoneNumber,
	}

	user, err := a.repository.GetByPhoneNumberIncludePassword(ctx, getByPhoneNumberInput)

	if err != nil {
		return nil, err
	}

	if user == nil {
		result.IsSuccess = false
		result.IsUserNotFound = true

		return result, nil
	}

	_, err = a.passwordAuth.CompareHashedPassword(user.Password, form.Password)

	if err != nil {

		result.IsSuccess = false
		result.IsUserNotFound = false

		return result, nil
	}

	expiredDuration, err := time.ParseDuration(os.Getenv("LOGIN_EXPIRATION_DURATION"))

	if err != nil {
		return nil, err
	}

	expiredTokenAt := time.Now().Add(expiredDuration)

	jwtToken, err := a.jwtAuth.GenerateJwt(modules.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APPLICATION_NAME"),
			ExpiresAt: jwt.NewNumericDate(expiredTokenAt),
		},
		UserId: user.Id,
	})

	if err != nil {
		return nil, err
	}

	user.LoginSuccessCount += 1

	updateUserInput := repository.UpdateUserInput{
		Id:                user.Id,
		PhoneNumber:       user.PhoneNumber,
		FullName:          user.FullName,
		LoginSuccessCount: user.LoginSuccessCount,
	}

	updateOutput, err := a.repository.Update(ctx, updateUserInput)

	if err != nil {
		return nil, err
	}

	if updateOutput.IsSuccessUpdate == false {
		return nil, err
	}

	result.IsSuccess = true
	result.IsUserNotFound = false
	result.HasValidationErrors = false
	result.ValidationErrors = map[string]string{}
	result.Credential = &AuthenticationCredential{
		Token: *jwtToken,
	}

	return result, nil
}

func (a AuthenticationService) Authorize(tokenString string) (*AuthorizationResult, error) {

	claims, err := a.jwtAuth.VerifyJwt(tokenString)

	if err != nil {
		return nil, err
	}

	result := &AuthorizationResult{
		IsAuthorized: true,
		UserId:       claims.UserId,
	}

	return result, nil
}

func NewAuthenticationService(repository repository.UserRepositoryInterface,
	passwordAuth modules.PasswordAuthInterface, jwtAuth modules.JsonWebTokenUtilInterface) AuthenticationServiceInterface {

	return AuthenticationService{
		repository:   repository,
		passwordAuth: passwordAuth,
		jwtAuth:      jwtAuth,
	}
}
