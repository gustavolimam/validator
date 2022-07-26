package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gustavolimam/validator/validations"
)

func New() *validator.Validate {
	validate := validator.New()

	// Register all custom validation we create
	validate.RegisterValidation("cpf", validations.CPFValidation)

	return validate
}
