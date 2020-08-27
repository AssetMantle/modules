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

type assetID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*assetID)(nil)

func (assetID assetID) Bytes() []byte {
	return append(
		assetID.ClassificationID.Bytes(),
		assetID.HashID.Bytes()...,
	)
}

func (assetID assetID) String() string {
	var values []string
	values = append(values, assetID.ClassificationID.String())
	values = append(values, assetID.HashID.String())
	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}

func (assetID assetID) Compare(id types.ID) int {
	return bytes.Compare(assetID.Bytes(), id.Bytes())
}

func readAssetID(assetIDString string) types.ID {
	idList := strings.Split(assetIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return assetID{
			ClassificationID: base.NewID(idList[0]),
			HashID:           base.NewID(idList[1]),
		}
	}
	return assetID{ClassificationID: base.NewID(""), HashID: base.NewID("")}
}

func assetIDFromInterface(id types.ID) assetID {
	switch value := id.(type) {
	case assetID:
		return value
	default:
		return assetIDFromInterface(readAssetID(id.String()))
	}
}
func generateKey(assetID types.ID) []byte {
	return append(StoreKeyPrefix, assetIDFromInterface(assetID).Bytes()...)
}

func NewAssetID(classificationID types.ID, immutables types.Immutables) types.ID {
	return assetID{
		ClassificationID: classificationID,
		HashID:           immutables.GetHashID(),
	}
}
