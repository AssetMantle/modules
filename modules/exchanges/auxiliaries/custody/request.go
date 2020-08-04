package swap

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type auxiliaryRequest struct {
	Order mappables.Order
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

func NewAuxiliaryRequest(order mappables.Order) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		Order: order,
	}
}
