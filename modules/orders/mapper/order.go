package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type order struct {
	ID         types.ID         `json:"id" valid:"required~required field id missing"`
	Immutables types.Immutables `json:"immutables" valid:"required field immutables missing"`
	Mutables   types.Mutables   `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.Order = (*order)(nil)

func (order order) GetID() types.ID {
	return order.ID
}

func (order order) GetChainID() types.ID {
	return orderIDFromInterface(order.ID).ChainID
}

func (order order) GetImmutables() types.Immutables {
	return order.Immutables
}

func (order order) GetMutables() types.Mutables {
	return order.Mutables
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
func NewOrder(orderID types.ID, mutables types.Mutables, immutables types.Immutables) mappables.Order {
	return order{
		ID:         orderID,
		Mutables:   mutables,
		Immutables: immutables,
	}
}
