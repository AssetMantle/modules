package entities

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Split interface {
	GetID() types.ID
	traits.Ownable
	traits.Transactional
}
