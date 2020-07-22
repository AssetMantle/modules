package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
)

type order struct {
	ID           schema.ID
	Burn         schema.Height
	Lock         schema.Height
	Immutables   schema.Immutables
	Mutables     schema.Mutables
	MakerAddress sdkTypes.AccAddress
	TakerAddress sdkTypes.AccAddress
	//SenderAddress       sdkTypes.AccAddress
	//FeeRecipientAddress sdkTypes.AccAddress
	MakerAssetAmount sdkTypes.Dec
	MakerAssetData   schema.ID
	//MakerFee            sdkTypes.Dec
	//MakerFeeAssetData   ID
	TakerAssetAmount sdkTypes.Dec
	TakerAssetData   schema.ID
	//TakerFee            sdkTypes.Dec
	//TakerFeeAssetData   ID
	//ExpirationTime      Height
	Salt schema.Height
}

var _ schema.Order = (*order)(nil)

func (order order) GetID() schema.ID {
	return order.ID
}

func (order order) GetChainID() schema.ID {
	return orderIDFromInterface(order.ID).ChainID
}

func (order order) GetClassificationID() schema.ID {
	return orderIDFromInterface(order.ID).ClassificationID
}

func (order order) GetBurn() schema.Height {
	return order.Burn
}

func (order order) CanBurn(currentHeight schema.Height) bool {
	return currentHeight.IsGreaterThan(order.Burn)
}

func (order order) GetLock() schema.Height {
	return order.Lock
}

func (order order) CanSend(currentHeight schema.Height) bool {
	return currentHeight.IsGreaterThan(order.Lock)
}

func (order order) GetImmutables() schema.Immutables {
	return order.Immutables
}

func (order order) GetMutables() schema.Mutables {
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
func (order order) GetMakerAssetData() schema.ID {
	return order.MakerAssetData
}

func (order order) GetTakerAssetAmount() sdkTypes.Dec {
	return order.TakerAssetAmount
}
func (order order) GetTakerAssetData() schema.ID {
	return order.TakerAssetData
}
func (order order) GetSalt() schema.Height {
	return order.Salt
}

func (order order) SetTakerAddress(takerAddress sdkTypes.AccAddress) schema.Order {
	order.TakerAddress = takerAddress
	return order

}

func NewOrder(orderID schema.ID, burn schema.Height, lock schema.Height, immutables schema.Immutables, mutables schema.Mutables,
	makerAddess sdkTypes.AccAddress, takerAddress sdkTypes.AccAddress,
	makerAssetAmount sdkTypes.Dec, makerAssetData schema.ID,
	takerAssetAmount sdkTypes.Dec, takerAssetData schema.ID, salt schema.Height) order {
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
