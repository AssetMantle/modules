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

type maintainerID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	IdentityID       types.ID `json:"identityID" valid:"required~required field identityID missing"`
}

var _ types.ID = (*maintainerID)(nil)

func (maintainerID maintainerID) Bytes() []byte {
	return append(
		maintainerID.ClassificationID.Bytes(),
		maintainerID.IdentityID.Bytes()...)

}

func (maintainerID maintainerID) String() string {
	var values []string
	values = append(values, maintainerID.ClassificationID.String())
	values = append(values, maintainerID.IdentityID.String())
	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}

func (maintainerID maintainerID) Equal(id types.ID) bool {
	return bytes.Compare(maintainerID.Bytes(), id.Bytes()) == 0
}

func readMaintainerID(maintainerIDString string) types.ID {
	idList := strings.Split(maintainerIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return maintainerID{
			ClassificationID: base.NewID(idList[0]),
			IdentityID:       base.NewID(idList[1]),
		}
	}
	return maintainerID{IdentityID: base.NewID(""), ClassificationID: base.NewID("")}
}

func maintainerIDFromInterface(id types.ID) maintainerID {
	switch value := id.(type) {
	case maintainerID:
		return value
	default:
		return maintainerIDFromInterface(readMaintainerID(id.String()))
	}
}
func generateKey(maintainerID types.ID) []byte {
	return append(StoreKeyPrefix, maintainerIDFromInterface(maintainerID).Bytes()...)
}
func NewMaintainerID(classificationID types.ID, identityID types.ID) types.ID {
	return maintainerID{
		ClassificationID: classificationID,
		IdentityID:       identityID,
	}
}
