package container

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type ValidationIoC struct {
	Validator *validator.Validate
}

func (v ValidationIoC) IsEmpty() bool {
	return (ValidationIoC{}) == v
}

func (v *ValidationIoC) Validate(model interface{}) error {
	return v.Validator.Struct(model)
}

var validation ValidationIoC
var validationSingleton sync.Once

func NewValidationService() ValidationIoC {
	validationSingleton.Do(func() {
		validator := validator.New()

		validation = ValidationIoC{
			Validator: validator,
		}
	})

	return validation
}
