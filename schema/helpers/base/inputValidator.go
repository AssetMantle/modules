package base

import (
	"regexp"
)

type InputValidator struct {
	expression string
}

func NewInputValidator(expression string) InputValidator {
	return InputValidator{expression}
}

func (inputValidator InputValidator) IsValid(values ...string) bool {
	for _, value := range values {
		valid, err := regexp.MatchString(inputValidator.expression, value)
		if !valid || err != nil {
			return false
		}
	}
	return true
}
