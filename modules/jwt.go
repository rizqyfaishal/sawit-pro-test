package modules

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserId      int64
}

type JsonWebTokenUtilInterface interface {
	GenerateJwt(claims CustomClaims) (*string, error)
	VerifyJwt(tokenString string) (*CustomClaims, error)
}
