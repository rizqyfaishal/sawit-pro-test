package handler

import (
	"github.com/SawitProRecruitment/UserService/services"
)

type Server struct {
	userService services.UserServiceInterface
	authenticationService services.AuthenticationServiceInterface
}

type NewServerOptions struct {
	UserService           services.UserServiceInterface
	AuthenticationService services.AuthenticationServiceInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		userService:           opts.UserService,
		authenticationService: opts.AuthenticationService,
	}
}
