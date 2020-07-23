package entities

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Classification interface {
	String() string
	GetID() types.ID
	GetTraits() types.Traits
}
