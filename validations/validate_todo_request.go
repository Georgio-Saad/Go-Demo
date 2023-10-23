package validations

import (
	"todogorest/data/request"

	"github.com/go-playground/validator/v10"
)

func ValidateTodoRequest(todoRequest request.TodoRequest) error {
	validate := validator.New()

	return validate.Struct(todoRequest)
}
