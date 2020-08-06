package reverse

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type auxiliaryRequest struct {
	MakerID      types.ID
	MakerSplit   sdkTypes.Dec
	MakerSplitID types.ID
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func auxiliaryRequestFromInterface(AuxiliaryRequest helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := AuxiliaryRequest.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(makerID types.ID, makerSplit sdkTypes.Dec, makerSplitID types.ID) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		MakerID:      makerID,
		MakerSplit:   makerSplit,
		MakerSplitID: makerSplitID,
	}
}
