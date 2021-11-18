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

type splitID struct {
	OwnerID   types.ID `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID types.ID `json:"ownableID" valid:"required~required field ownableID missing"`
}

var _ types.ID = (*splitID)(nil)

func (splitID splitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)

}

func (splitID splitID) String() string {
	var values []string
	values = append(values, splitID.OwnerID.String())
	values = append(values, splitID.OwnableID.String())
	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}

func (splitID splitID) Equal(id types.ID) bool {
	return bytes.Compare(splitID.Bytes(), id.Bytes()) == 0
}

func readSplitID(splitIDString string) types.ID {
	idList := strings.Split(splitIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return splitID{
			OwnerID:   base.NewID(idList[0]),
			OwnableID: base.NewID(idList[1]),
		}
	} else if len(idList) == 3 {
		if idList[0] == "orders" {
			return splitID{
				OwnerID:   base.NewID(idList[0]),
				OwnableID: base.NewID(strings.Join([]string{idList[1], idList[2]}, constants.FirstOrderCompositeIDSeparator)),
			}
		} else {
			return splitID{
				OwnerID:   base.NewID(strings.Join([]string{idList[0], idList[1]}, constants.FirstOrderCompositeIDSeparator)),
				OwnableID: base.NewID(idList[2]),
			}
		}
	} else if len(idList) == 4 {
		return splitID{
			OwnerID:   base.NewID(strings.Join([]string{idList[0], idList[1]}, constants.FirstOrderCompositeIDSeparator)),
			OwnableID: base.NewID(strings.Join([]string{idList[2], idList[3]}, constants.FirstOrderCompositeIDSeparator)),
		}
	}
	return splitID{OwnerID: base.NewID(""), OwnableID: base.NewID("")}
}

func splitIDFromInterface(id types.ID) splitID {
	switch value := id.(type) {
	case splitID:
		return value
	default:
		return splitIDFromInterface(readSplitID(id.String()))
	}
}
func generateKey(splitID types.ID) []byte {
	return append(StoreKeyPrefix, splitIDFromInterface(splitID).Bytes()...)
}
func NewSplitID(ownerID types.ID, ownableID types.ID) types.ID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
