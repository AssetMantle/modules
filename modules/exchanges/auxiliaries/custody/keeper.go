package swap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper     helpers.Mapper
	BankKeeper bank.Keeper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) error {
	//auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)

	return nil
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.AuxiliaryKeeper {
	auxiliaryKeeper := auxiliaryKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case bank.Keeper:
			auxiliaryKeeper.BankKeeper = value
		}
	}
	return auxiliaryKeeper
}
