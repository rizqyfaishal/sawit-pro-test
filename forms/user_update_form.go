package forms

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserUpdateForm struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"  validate:"omitempty,min=10,max=13,startswith=+62"`
	FullName    string `form:"full_name" json:"full_name" validate:"omitempty,min=3,max=60"`
}

func (u UserUpdateForm) TranslateField(field string) string {

	switch field {

	case "PhoneNumber":
		return "Phone number"
	case "FullName":
		return "Full name"
	}

	return "unknown"
}

func (u UserUpdateForm) GetErrorMessage(fieldError validator.FieldError) string {

	translatedField := u.TranslateField(fieldError.Field())

	switch fieldError.Tag() {
	case "min":
		return fmt.Sprintf("%s must have minimum %s characters long", translatedField, fieldError.Param())
	case "required":
		return fmt.Sprintf("%s is required", translatedField)
	case "max":
		return fmt.Sprintf("%s must have maximum %s characters long", translatedField, fieldError.Param())
	case "startswith":
		return fmt.Sprintf("%s must starts with %s", translatedField, fieldError.Param())
	}

	return "unknown error"
}

func (u UserUpdateForm) GetFormField(fieldError validator.FieldError) string {

	switch fieldError.Field() {

	case "PhoneNumber":
		return "phone_number"
	case "FullName":
		return "full_name"
	}

	return "unknown"
}
