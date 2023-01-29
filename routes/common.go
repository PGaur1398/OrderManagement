package routes

import (
	"fmt"

	"github.com/go-playground/validator"
)

/*
	The function iterates through the errors in the validator.ValidationErrors struct and

creates a custom error message string with the field name and type of error for each validation error.
*/
func StructValidatorErrorHandling(err error) string {
	var customErrorMsg string
	for _, err := range err.(validator.ValidationErrors) {
		fieldErrMsg := "Field validation for '%s' failed with type: '%s'"
		if customErrorMsg == "" {
			customErrorMsg += fmt.Sprintf(fieldErrMsg, err.Field(), err.Type())
		} else {
			customErrorMsg += ", " + fmt.Sprintf(fieldErrMsg, err.Field(), err.Type())
		}
	}
	return customErrorMsg
}
