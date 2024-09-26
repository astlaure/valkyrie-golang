package core

import (
	"github.com/go-playground/validator/v10"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}

	FormErrors struct {
		Messages map[string]string
	}
)

func (e *FormErrors) Error() string {
	return "Form Validation errors"
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		v := make(map[string]string)

		for _, element := range errs {
			v[element.Field()] = element.Tag()
		}

		return &FormErrors{Messages: v}
	}
	return nil
}

var Validator = &CustomValidator{validator: validator.New()}
