package forms

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserLoginForm struct {
	PhoneNumber string `form:"phone_number" json:"phone_number" validate:"required"`
	Password    string `form:"password" json:"password" validate:"required"`
}

func (u UserLoginForm) GetFormField(fieldError validator.FieldError) string {

	switch fieldError.Field() {

	case "PhoneNumber":
		return "phone_number"
	case "Password":
		return "password"
	}

	return "unknown"
}

func (u UserLoginForm) TranslateField(field string) string {

	switch field {

	case "PhoneNumber":
		return "Phone number"
	case "Password":
		return "Password"
	}

	return "unknown"
}

func (u UserLoginForm) GetErrorMessage(fieldError validator.FieldError) string {

	translatedField := u.TranslateField(fieldError.Field())

	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", translatedField)
	}

	return "unknown error"
}
