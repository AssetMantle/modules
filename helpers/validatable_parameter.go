package helpers

import (
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/parameters"
)

type ValidatableParameter interface {
	GetParameter() parameters.Parameter
	Mutate(data.Data) ValidatableParameter
	Validate() error
}
