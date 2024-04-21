package forms

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/validators"
	"github.com/go-playground/validator/v10"
)

type UserRegisterForm struct {
	PhoneNumber string `form:"phone_number" json:"phone_number" validate:"required,min=10,max=13,startswith=+62"`
	FullName    string `form:"full_name" json:"full_name" validate:"required,min=3,max=60"`
	Password    string `form:"password" json:"password" validate:"required,min=6,max=64,atl_x_capital_char=1,atl_x_special_char=1"`
}

func (c UserRegisterForm) GetFormField(fieldError validator.FieldError) string {

	switch fieldError.Field() {

	case "PhoneNumber":
		return "phone_number"
	case "FullName":
		return "full_name"
	case "Password":
		return "password"
	}

	return "unknown"
}

func (c UserRegisterForm) TranslateField(field string) string {

	switch field {

	case "PhoneNumber":
		return "Phone number"
	case "FullName":
		return "Full name"
	case "Password":
		return "Password"
	}

	return "unknown"
}

func (c UserRegisterForm) GetErrorMessage(fieldError validator.FieldError) string {

	translatedField := c.TranslateField(fieldError.Field())

	switch fieldError.Tag() {
	case "min":
		return fmt.Sprintf("%s must have minimum %s characters long", translatedField, fieldError.Param())
	case "required":
		return fmt.Sprintf("%s is required", translatedField)
	case "max":
		return fmt.Sprintf("%s must have maximum %s characters long", translatedField, fieldError.Param())
	case "startswith":
		return fmt.Sprintf("%s must starts with %s", translatedField, fieldError.Param())
	case validators.AtLeastXCapitalCharValidationTag:
		return fmt.Sprintf("%s must contains at least %s captial characters", translatedField, fieldError.Param())
	case validators.AtLeastXSpecialCharValidationTag:
		return fmt.Sprintf("%s must contains at least %s special characters", translatedField, fieldError.Param())
	}

	return "unknown error"
}
