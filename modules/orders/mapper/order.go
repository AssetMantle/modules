package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order struct {
	ID               types.ID
	Burn             types.Height
	Lock             types.Height
	Immutables       types.Immutables
	MakerAddress     sdkTypes.AccAddress
	TakerAddress     sdkTypes.AccAddress
	MakerAssetAmount sdkTypes.Dec
	MakerAssetData   traits.Exchangeable
	TakerAssetAmount sdkTypes.Dec
	TakerAssetData   traits.Exchangeable
	Salt             types.Height
}

var _ mappables.Order = (*Order)(nil)

func (order Order) GetID() types.ID {
	return order.ID
}

func (order Order) GetChainID() types.ID {
	return orderIDFromInterface(order.ID).ChainID
}

func (order Order) GetBurn() types.Height {
	return order.Burn
}

func (order Order) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Burn)
}

func (order Order) GetLock() types.Height {
	return order.Lock
}

func (order Order) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Lock)
}

func (order Order) GetImmutables() types.Immutables {
	return order.Immutables
}
func (order Order) GetMakerAddress() sdkTypes.AccAddress {
	return order.MakerAddress
}
func (order Order) GetTakerAddress() sdkTypes.AccAddress {
	return order.TakerAddress
}

func (order Order) GetMakerAssetAmount() sdkTypes.Dec {
	return order.MakerAssetAmount
}
func (order Order) GetMakerAssetData() traits.Exchangeable {
	return order.MakerAssetData
}

func (order Order) GetTakerAssetAmount() sdkTypes.Dec {
	return order.TakerAssetAmount
}
func (order Order) GetTakerAssetData() traits.Exchangeable {
	return order.TakerAssetData
}
func (order Order) GetSalt() types.Height {
	return order.Salt
}
func (order Order) SetTakerAddress(takerAddress sdkTypes.AccAddress) mappables.Order {
	order.TakerAddress = takerAddress
	return order
}
func (order Order) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(order)
}
func (order Order) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &order)
	return order
}
func orderPrototype() traits.Mappable {
	return Order{}
}
func NewOrder(orderID types.ID, burn types.Height, lock types.Height, immutables types.Immutables,
	makerAddress sdkTypes.AccAddress, takerAddress sdkTypes.AccAddress, makerAssetAmount sdkTypes.Dec,
	makerAssetData traits.Exchangeable, takerAssetAmount sdkTypes.Dec, takerAssetData traits.Exchangeable, salt types.Height) mappables.Order {
	return Order{
		ID:               orderID,
		Burn:             burn,
		Lock:             lock,
		Immutables:       immutables,
		MakerAddress:     makerAddress,
		TakerAddress:     takerAddress,
		MakerAssetAmount: makerAssetAmount,
		MakerAssetData:   makerAssetData,
		TakerAssetAmount: takerAssetAmount,
		TakerAssetData:   takerAssetData,
		Salt:             salt,
	}
}
