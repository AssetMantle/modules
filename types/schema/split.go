package schema

import (
	"github.com/persistenceOne/persistenceSDK/types/trait"
)

type Split interface {
	GetID() ID
	trait.Ownable
	trait.Transactional
}
