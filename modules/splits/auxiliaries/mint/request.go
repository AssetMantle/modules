package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type auxiliaryRequest struct {
	OwnerID   schema.ID    `json:"OwnerId" valid:"required~Enter the OwnerID"`
	OwnableID schema.ID    `json:"OwnableId" valid:"required~Enter the OwnableID"`
	Split     sdkTypes.Dec `json:"split" valid:"required~Enter the Split,matches(^[0-9]$)~Split is Invalid"`
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

func NewAuxiliaryRequest(ownerID schema.ID, ownableID schema.ID, split sdkTypes.Dec) utility.AuxiliaryRequest {
	return &auxiliaryRequest{
		OwnerID:   ownerID,
		OwnableID: ownableID,
		Split:     split,
	}
}
