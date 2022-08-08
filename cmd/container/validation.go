package container

import (
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
		validator.RegisterCustomTypeFunc(ValidateUUID, uuid.UUID{})

		validation = ValidationIoC{
			Validator: validator,
		}
	})

	return validation
}

func ValidateUUID(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		return valuer.String()
	}
	return nil
}
