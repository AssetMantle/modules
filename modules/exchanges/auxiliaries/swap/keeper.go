package swap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) error {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	order := auxiliaryRequest.Order
	makeOrderCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(order.GetMakerAssetData().String(), order.GetMakerAssetAmount().TruncateInt()))
	takeOrderCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(order.GetTakerAssetData().String(), order.GetTakerAssetAmount().TruncateInt()))
	error := auxiliaryRequest.BankKeeper.SendCoins(context, order.GetMakerAddress(), order.GetTakerAddress(), makeOrderCoins)
	if error != nil {
		return error
	}
	error = auxiliaryRequest.BankKeeper.SendCoins(context, order.GetTakerAddress(), order.GetMakerAddress(), takeOrderCoins)
	if error != nil {
		return error
	}

	return nil
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
