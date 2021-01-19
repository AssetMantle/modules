/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
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

func (orderID orderID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, orderID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, orderID.MakerOwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.TakerOwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.MakerID.Bytes()...)
	Bytes = append(Bytes, orderID.HashID.Bytes()...)

	return Bytes
}
func (orderID orderID) String() string {
	var values []string
	values = append(values, orderID.ClassificationID.String())
	values = append(values, orderID.MakerOwnableID.String())
	values = append(values, orderID.TakerOwnableID.String())
	values = append(values, orderID.MakerID.String())
	values = append(values, orderID.HashID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (orderID orderID) Equals(id types.ID) bool {
	return bytes.Equal(orderID.Bytes(), id.Bytes())
}
func (orderID orderID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(orderID.Bytes())
}
func (orderID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, orderID{})
}
func (orderID orderID) IsPartial() bool {
	return len(orderID.HashID.Bytes()) == 0
}
func (orderID orderID) Matches(key helpers.Key) bool {
	return orderID.Equals(orderIDFromInterface(key))
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
