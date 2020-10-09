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

type identityID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*identityID)(nil)

func (identityID identityID) Bytes() []byte {
	return append(
		identityID.ClassificationID.Bytes(),
		identityID.HashID.Bytes()...,
	)
}

func (identityID identityID) String() string {
	var values []string
	values = append(values, identityID.ClassificationID.String())
	values = append(values, identityID.HashID.String())
	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}

func (identityID identityID) Equals(id types.ID) bool {
	return bytes.Compare(identityID.Bytes(), id.Bytes()) == 0
}

func readIdentityID(identityIDString string) types.ID {
	idList := strings.Split(identityIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return identityID{
			ClassificationID: base.NewID(idList[0]),
			HashID:           base.NewID(idList[1]),
		}
	}
	return identityID{ClassificationID: base.NewID(""), HashID: base.NewID("")}
}

func identityIDFromInterface(id types.ID) identityID {
	switch value := id.(type) {
	case identityID:
		return value
	default:
		return identityIDFromInterface(readIdentityID(id.String()))
	}
}
func generateKey(identityID types.ID) []byte {
	return append(StoreKeyPrefix, identityIDFromInterface(identityID).Bytes()...)
}
func NewIdentityID(classificationID types.ID, hashID types.ID) types.ID {
	return identityID{
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
