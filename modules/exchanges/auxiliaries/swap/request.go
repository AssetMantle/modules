package swap

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type auxiliaryRequest struct {
	Order      mappables.Order
	BankKeeper bank.Keeper
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

func NewAuxiliaryRequest(order mappables.Order, bankKeeper bank.Keeper) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		Order:      order,
		BankKeeper: bankKeeper,
	}
}
