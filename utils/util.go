package utils

import (
	"github.com/SawitProRecruitment/UserService/forms"
	"github.com/go-playground/validator/v10"
)

func CollectValidationErrorMessages(parseErrorMessage forms.ParseableErrorMessageInterface, errors []validator.FieldError) map[string]string {

	errorMessages := make(map[string]string)

	for _, err := range errors {

		errorMessage := parseErrorMessage.GetErrorMessage(err)
		errorMessages[parseErrorMessage.GetFormField(err)] = errorMessage
	}

	return errorMessages
}
