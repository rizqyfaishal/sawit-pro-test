package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
)

type PasswordAdditionalValidator struct{}

var CapitalCharacterRegex = regexp.MustCompile("[A-Z]")
var NonAlphanumericCharacterRegex = regexp.MustCompile("[^a-zA-Z\\d\\s:]")

const AtLeastXCapitalCharValidationTag = "atl_x_capital_char"
const AtLeastXSpecialCharValidationTag = "atl_x_special_char"

func AtLeastXCapitalCharValidation(fl validator.FieldLevel) bool {

	minimumNumber, err := strconv.Atoi(fl.Param())

	if err != nil {
		return false
	}

	return validateAtLeastXCapitalChar(fl.Field().String(), minimumNumber)
}

func AtLeastXSpecialCharValidation(fl validator.FieldLevel) bool {
	minimumNumber, err := strconv.Atoi(fl.Param())

	if err != nil {
		return false
	}

	return validateAtLeastXSpecialChar(fl.Field().String(), minimumNumber)
}

func validateAtLeastXSpecialChar(str string, minimumNumber int) bool {

	return len(NonAlphanumericCharacterRegex.FindStringSubmatch(str)) >= minimumNumber
}

func validateAtLeastXCapitalChar(str string, minimumNumber int) bool {

	return len(CapitalCharacterRegex.FindStringSubmatch(str)) >= minimumNumber
}
