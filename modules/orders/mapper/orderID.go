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
	OwnableID     types.ID `json:"ownableID"`
	OwnablePairID types.ID `json:"ownablePairID"`
}

var _ types.ID = (*orderID)(nil)

func (orderID orderID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, orderID.OwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.OwnablePairID.Bytes()...)
	return Bytes
}

func (orderID orderID) String() string {
	var values []string
	values = append(values, orderID.OwnableID.String())
	values = append(values, orderID.OwnablePairID.String())
	return strings.Join(values, constants.CompositeIDSeparator)
}

func (orderID orderID) Compare(id types.ID) int {
	return bytes.Compare(orderID.Bytes(), id.Bytes())
}

func readOrderID(orderIDString string) types.ID {
	idList := strings.Split(orderIDString, constants.IDSeparator)
	if len(idList) == 3 {
		return orderID{
			ChainID:       base.NewID(idList[0]),
			MaintainersID: base.NewID(idList[1]),
			HashID:        base.NewID(idList[2]),
		}
	}
	return orderID{ChainID: base.NewID(""), MaintainersID: base.NewID(""), HashID: base.NewID("")}
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
func NewOrderID(chainID types.ID, maintainersID types.ID, hashID types.ID) types.ID {
	return orderID{
		ChainID:       chainID,
		MaintainersID: maintainersID,
		HashID:        hashID,
	}
}
