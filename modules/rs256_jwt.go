package modules

import (
	"fmt"
)
import "github.com/golang-jwt/jwt/v5"

type RS256Jwt struct {
	privateKey []byte
	publicKey  []byte
}

func (r RS256Jwt) GenerateJwt(claims CustomClaims) (*string, error) {

	signedKey, err := jwt.ParseRSAPrivateKeyFromPEM(r.privateKey)

	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signedKey)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (r RS256Jwt) VerifyJwt(tokenString string) (*CustomClaims, error) {

	signedPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(r.publicKey)

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {

		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {

			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return signedPublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	userIdFloat := claims["UserId"].(float64)

	customClaims := CustomClaims{
		UserId: int64(userIdFloat),
	}

	return &customClaims, nil
}

func NewRS256Jwt(privateKey []byte, publicKey []byte) JsonWebTokenUtilInterface {
	return RS256Jwt{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}
