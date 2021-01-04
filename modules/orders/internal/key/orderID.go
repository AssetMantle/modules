/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strings"
)

type orderID struct {
	ClassificationID types.ID `json:"classificationID"`
	MakerOwnableID   types.ID `json:"makerOwnableID"`
	TakerOwnableID   types.ID `json:"takerOwnableID"`
	MakerID          types.ID `json:"makerID"`
	HashID           types.ID `json:"hashID"`
}

var _ types.ID = (*orderID)(nil)
var _ helpers.Key = (*orderID)(nil)

func (OrderID orderID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, OrderID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, OrderID.MakerOwnableID.Bytes()...)
	Bytes = append(Bytes, OrderID.TakerOwnableID.Bytes()...)
	Bytes = append(Bytes, OrderID.MakerID.Bytes()...)
	Bytes = append(Bytes, OrderID.HashID.Bytes()...)
	return Bytes
}
func (OrderID orderID) String() string {
	var values []string
	values = append(values, OrderID.ClassificationID.String())
	values = append(values, OrderID.MakerOwnableID.String())
	values = append(values, OrderID.TakerOwnableID.String())
	values = append(values, OrderID.MakerID.String())
	values = append(values, OrderID.HashID.String())
	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (OrderID orderID) Equals(id types.ID) bool {
	return bytes.Compare(OrderID.Bytes(), id.Bytes()) == 0
}
func (OrderID orderID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(OrderID.Bytes())
}
func (orderID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, orderID{})
}
func (OrderID orderID) IsPartial() bool {
	if len(OrderID.HashID.Bytes()) > 0 {
		return false
	}
	return true
}
func (OrderID orderID) Matches(key helpers.Key) bool {
	switch value := key.(type) {
	case orderID:
		return bytes.Compare(OrderID.Bytes(), value.Bytes()) == 0
	default:
		return false
	}
}

func New(id types.ID) helpers.Key {
	return orderIDFromInterface(id)
}

func NewOrderID(classificationID types.ID, makerOwnableID types.ID, takerOwnableID types.ID, makerID types.ID, immutables types.Immutables) types.ID {
	return orderID{
		ClassificationID: classificationID,
		MakerOwnableID:   makerOwnableID,
		TakerOwnableID:   takerOwnableID,
		MakerID:          makerID,
		HashID:           immutables.GetHashID(),
	}
}
