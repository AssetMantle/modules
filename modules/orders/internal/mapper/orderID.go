/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
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

func (orderID orderID) Compare(id types.ID) int {
	return bytes.Compare(orderID.Bytes(), id.Bytes())
}

func readOrderID(orderIDString string) types.ID {
	idList := strings.Split(orderIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 5 {
		return orderID{
			ClassificationID: base.NewID(idList[0]),
			MakerOwnableID:   base.NewID(idList[1]),
			TakerOwnableID:   base.NewID(idList[2]),
			MakerID:          base.NewID(idList[3]),
			HashID:           base.NewID(idList[4]),
		}
	}
	return orderID{ClassificationID: base.NewID(""), MakerOwnableID: base.NewID(""), TakerOwnableID: base.NewID(""), MakerID: base.NewID(""), HashID: base.NewID("")}
}

func orderIDFromInterface(id types.ID) orderID {
	switch value := id.(type) {
	case orderID:
		return value
	default:
		return orderIDFromInterface(readOrderID(id.String()))
	}
}
func generateKey(orderID types.ID) []byte {
	return append(StoreKeyPrefix, orderIDFromInterface(orderID).Bytes()...)
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
