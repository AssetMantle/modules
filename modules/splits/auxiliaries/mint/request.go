package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type auxiliaryRequest struct {
	OwnerID   schema.ID
	OwnableID schema.ID
	Split     sdkTypes.Dec
}

var _ utility.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func auxiliaryRequestFromInterface(AuxiliaryRequest utility.AuxiliaryRequest) auxiliaryRequest {
	switch value := AuxiliaryRequest.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(toID schema.ID, split sdkTypes.Dec) utility.AuxiliaryRequest {
	return auxiliaryRequest{
		ToID:  toID,
		Split: split,
	}
}
