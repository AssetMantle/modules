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
	if Error := auxiliaryRequest.BankKeeper.SendCoins(context, order.GetMakerAddress(), order.GetTakerAddress(), makeOrderCoins); Error != nil {
		return Error
	}
	if Error := auxiliaryRequest.BankKeeper.SendCoins(context, order.GetTakerAddress(), order.GetMakerAddress(), takeOrderCoins); Error != nil {
		return Error
	}
	return nil
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
