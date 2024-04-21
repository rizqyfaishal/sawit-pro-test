package handler

import (
	"github.com/SawitProRecruitment/UserService/consts"
	"github.com/SawitProRecruitment/UserService/forms"
	"github.com/SawitProRecruitment/UserService/responses"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Register a user
// (POST /users/register)
func (s *Server) Register(ctx echo.Context) error {

	var userRegisterForm forms.UserRegisterForm

	if err := ctx.Bind(&userRegisterForm); err != nil {

		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	registerResult, err := s.userService.Register(userRegisterForm)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	if registerResult.HasValidationErrors {
		return ctx.JSON(http.StatusBadRequest, registerResult.ValidationErrors)
	}

	return ctx.JSON(http.StatusOK, registerResult.User)
}

// Update user profile
// (PUT /users)
func (s *Server) UpdateUser(ctx echo.Context) error {

	authorizedUserId := ctx.Get(consts.ContextAuthorizedUsedId).(int64)

	var updateUserForm forms.UserUpdateForm

	if err := ctx.Bind(&updateUserForm); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	if utils.StringIsEmpty(updateUserForm.PhoneNumber) == false {
		user, err := s.userService.GetByPhoneNumber(updateUserForm.PhoneNumber)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		}

		if user != nil {
			return ctx.JSON(http.StatusConflict, responses.BadRequestResponse{
				ErrorMessage: "Conflicted",
			})
		}
	}

	updateResult, err := s.userService.Update(authorizedUserId, updateUserForm)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	if updateResult.HasValidationErrors {
		return ctx.JSON(http.StatusBadRequest, updateResult.ValidationErrors)
	}

	return ctx.JSON(http.StatusOK, updateResult.User)
}

// Login
// (POST /users/login)
func (s *Server) Login(ctx echo.Context) error {
	var userLoginForm forms.UserLoginForm

	if err := ctx.Bind(&userLoginForm); err != nil {

		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	authenticationResult, err := s.authenticationService.Authenticate(userLoginForm)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	if authenticationResult.HasValidationErrors {
		return ctx.JSON(http.StatusBadRequest, authenticationResult.ValidationErrors)
	}

	if authenticationResult.IsUserNotFound || authenticationResult.IsSuccess == false {
		badRequestResponse := responses.BadRequestResponse{
			ErrorMessage: "Login failed. Please enter correct phone number and password.",
		}
		return ctx.JSON(http.StatusBadRequest, badRequestResponse)
	}

	return ctx.JSON(http.StatusOK, authenticationResult.Credential)
}

// Get User Profile
// (GET /users/me)
func (s *Server) GetMyProfile(ctx echo.Context) error {

	authorizedUserId := ctx.Get(consts.ContextAuthorizedUsedId).(int64)

	user, err := s.userService.GetById(authorizedUserId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	if user == nil {
		return ctx.JSON(http.StatusNotFound, "User not found")
	}

	return ctx.JSON(http.StatusOK, user)
}
