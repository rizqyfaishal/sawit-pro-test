package forms

import "github.com/go-playground/validator/v10"

type ParseableErrorMessageInterface interface {
	GetErrorMessage(fieldError validator.FieldError) string
	GetFormField(fieldError validator.FieldError) string
}
