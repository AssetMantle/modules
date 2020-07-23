package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type order struct {
	ID               types.ID
	Burn             types.Height
	Lock             types.Height
	Immutables       types.Immutables
	Mutables         types.Mutables
	MakerAddress     sdkTypes.AccAddress
	TakerAddress     sdkTypes.AccAddress
	MakerAssetAmount sdkTypes.Dec
	MakerAssetData   types.ID
	TakerAssetAmount sdkTypes.Dec
	TakerAssetData   types.ID
	Salt             types.Height
}

var _ entities.Order = (*order)(nil)

func (order order) GetID() types.ID {
	return order.ID
}

func (order order) GetChainID() types.ID {
	return orderIDFromInterface(order.ID).ChainID
}

func (order order) GetClassificationID() types.ID {
	return orderIDFromInterface(order.ID).ClassificationID
}

func (order order) GetBurn() types.Height {
	return order.Burn
}

func (order order) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Burn)
}

func (order order) GetLock() types.Height {
	return order.Lock
}

func (order order) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Lock)
}

func (order order) GetImmutables() types.Immutables {
	return order.Immutables
}

func (order order) GetMutables() types.Mutables {
	return order.Mutables
}

func (order order) GetMakerAddress() sdkTypes.AccAddress {
	return order.MakerAddress
}
func (order order) GetTakerAddress() sdkTypes.AccAddress {
	return order.TakerAddress
}

func (order order) GetMakerAssetAmount() sdkTypes.Dec {
	return order.MakerAssetAmount
}
func (order order) GetMakerAssetData() types.ID {
	return order.MakerAssetData
}

func (order order) GetTakerAssetAmount() sdkTypes.Dec {
	return order.TakerAssetAmount
}
func (order order) GetTakerAssetData() types.ID {
	return order.TakerAssetData
}
func (order order) GetSalt() types.Height {
	return order.Salt
}

func (order order) SetTakerAddress(takerAddress sdkTypes.AccAddress) entities.Order {
	order.TakerAddress = takerAddress
	return order

}

func NewOrder(orderID types.ID, burn types.Height, lock types.Height, immutables types.Immutables, mutables types.Mutables,
	makerAddess sdkTypes.AccAddress, takerAddress sdkTypes.AccAddress,
	makerAssetAmount sdkTypes.Dec, makerAssetData types.ID,
	takerAssetAmount sdkTypes.Dec, takerAssetData types.ID, salt types.Height) order {
	return order{
		ID:               orderID,
		Burn:             burn,
		Lock:             lock,
		Immutables:       immutables,
		Mutables:         mutables,
		MakerAddress:     makerAddess,
		TakerAddress:     takerAddress,
		MakerAssetAmount: makerAssetAmount,
		MakerAssetData:   makerAssetData,
		TakerAssetAmount: takerAssetAmount,
		TakerAssetData:   takerAssetData,
		Salt:             salt,
	}
}

//func NewOrder(makerAddess sdkTypes.AccAddress, takerAddress sdkTypes.AccAddress, senderAddress sdkTypes.AccAddress, feeRecipientAddress sdkTypes.AccAddress,
//	makerAssetAmount sdkTypes.Dec, makerAssetData ID, makerFee sdkTypes.Dec, makerFeeAssetData ID,
//	takerAssetAmount sdkTypes.Dec, takerAssetData ID, takerFee sdkTypes.Dec, takerFeeAssetData ID,
//	expirationTime Height, salt Height) Order {
//	return order{
//		MakerAddress:        makerAddess,
//		TakerAddress:        takerAddress,
//		SenderAddress:       senderAddress,
//		FeeRecipientAddress: feeRecipientAddress,
//		MakerAssetAmount:    makerAssetAmount,
//		MakerAssetData:      makerAssetData,
//		MakerFee:            makerFee,
//		MakerFeeAssetData:   makerFeeAssetData,
//		TakerAssetAmount:    takerAssetAmount,
//		TakerAssetData:      takerAssetData,
//		TakerFee:            takerFee,
//		TakerFeeAssetData:   takerFeeAssetData,
//		ExpirationTime:      expirationTime,
//		Salt:                salt,
//	}
//}
