package services

type Services struct {
	Authentication AuthenticationServiceInterface
	User           UserServiceInterface
}
