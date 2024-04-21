package services

import (
	"github.com/SawitProRecruitment/UserService/pojos"
)

type UpdateResult struct {
	User                pojos.User
	HasValidationErrors bool
	ValidationErrors    map[string]string
}

type RegisterResult struct {
	User                pojos.User
	ValidationErrors    map[string]string
	HasValidationErrors bool
}

type AuthenticationCredential struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

type AuthenticationResult struct {
	IsSuccess      bool
	IsUserNotFound bool

	HasValidationErrors bool
	ValidationErrors    map[string]string
	Credential          *AuthenticationCredential
}

type AuthorizationResult struct {
	IsAuthorized bool
	UserId       int64
}
