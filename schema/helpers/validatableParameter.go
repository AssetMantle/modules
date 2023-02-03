package helpers

import (
	"github.com/AssetMantle/modules/schema/data"
)

type ValidatableParameter interface {
	GetParameter() Parameter
	Mutate(data.AnyData) ValidatableParameter
	GetValidator() func(i interface{}) error
	Validate() error
}
