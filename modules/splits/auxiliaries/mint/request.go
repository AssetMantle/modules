package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type auxiliaryRequest struct {
	OwnerID   types.ID     `json:"OwnerId" valid:"required~Enter the OwnerID"`
	OwnableID types.ID     `json:"OwnableId" valid:"required~Enter the OwnableID"`
	Split     sdkTypes.Dec `json:"split" valid:"required~Enter the Split,matches(^[0-9]$)~Split is Invalid"`
}

var _ utilities.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func auxiliaryRequestFromInterface(AuxiliaryRequest utilities.AuxiliaryRequest) auxiliaryRequest {
	switch value := AuxiliaryRequest.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(ownerID types.ID, ownableID types.ID, split sdkTypes.Dec) utilities.AuxiliaryRequest {
	return &auxiliaryRequest{
		OwnerID:   ownerID,
		OwnableID: ownableID,
		Split:     split,
	}
}
