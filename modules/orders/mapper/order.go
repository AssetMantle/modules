package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type order struct {
	ID               types.ID
	Burn             types.Height
	Lock             types.Height
	Immutables       types.Immutables
	MakerID          types.ID
	TakerID          types.ID
	MakerAssetAmount sdkTypes.Dec
	MakerAssetData   types.ID
	TakerAssetAmount sdkTypes.Dec
	TakerAssetData   types.ID
	Salt             types.Height
}

var _ mappables.Order = (*order)(nil)

func (order order) GetID() types.ID {
	return order.ID
}

func (order order) GetChainID() types.ID {
	return orderIDFromInterface(order.ID).ChainID
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
func (order order) GetMakerID() types.ID {
	return order.MakerID
}
func (order order) GetTakerID() types.ID {
	return order.TakerID
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
func (order order) SetTakerID(takerAddress types.ID) mappables.Order {
	order.TakerID = takerAddress
	return order
}
func (order order) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(order)
}
func (order order) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &order)
	return order
}
func orderPrototype() traits.Mappable {
	return order{}
}
func NewOrder(orderID types.ID, burn types.Height, lock types.Height, immutables types.Immutables,
	makerID types.ID, takerID types.ID, makerAssetAmount sdkTypes.Dec, makerAssetData types.ID,
	takerAssetAmount sdkTypes.Dec, takerAssetData types.ID, salt types.Height) mappables.Order {
	return order{
		ID:               orderID,
		Burn:             burn,
		Lock:             lock,
		Immutables:       immutables,
		MakerID:          makerID,
		TakerID:          takerID,
		MakerAssetAmount: makerAssetAmount,
		MakerAssetData:   makerAssetData,
		TakerAssetAmount: takerAssetAmount,
		TakerAssetData:   takerAssetData,
		Salt:             salt,
	}
}
