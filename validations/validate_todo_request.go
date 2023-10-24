package validations

import (
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(s interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(s)
}
