package services

import (
	"github.com/SawitProRecruitment/UserService/forms"
	"github.com/SawitProRecruitment/UserService/pojos"
)

type UserServiceInterface interface {
	Register(form forms.UserRegisterForm) (*RegisterResult, error)
	Update(userId int64, form forms.UserUpdateForm) (*UpdateResult, error)
	GetById(userId int64) (*pojos.User, error)
	GetByPhoneNumber(phoneNumber string) (*pojos.User, error)
}

type AuthenticationServiceInterface interface {
	Authenticate(form forms.UserLoginForm) (*AuthenticationResult, error)
	Authorize(token string) (*AuthorizationResult, error)
}
