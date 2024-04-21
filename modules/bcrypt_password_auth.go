package modules

import "golang.org/x/crypto/bcrypt"

const GenerateHashedPasswordCost int = 10

type BcryptPasswordAuth struct{}

func (b BcryptPasswordAuth) GenerateHashedPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), GenerateHashedPasswordCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (b BcryptPasswordAuth) CompareHashedPassword(hashedPassword, password string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}
