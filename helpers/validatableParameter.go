package helpers

import (
	"github.com/AssetMantle/schema/x/data"
	"github.com/AssetMantle/schema/x/parameters"
)

type ValidatableParameter interface {
	GetParameter() parameters.Parameter
	Mutate(data.Data) ValidatableParameter
	GetValidator() func(i interface{}) error
	Validate() error
}
