/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type assetID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

func (AssetID assetID) GenerateStoreKeyBytes() []byte {
	return append(storeKeyPrefix, AssetID.Bytes()...)
}

var _ types.ID = (*assetID)(nil)
var _ helpers.Key = (*assetID)(nil)

func (AssetID assetID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, AssetID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, AssetID.HashID.Bytes()...)
	return Bytes
}

func (AssetID assetID) String() string {
	var values []string
	values = append(values, AssetID.ClassificationID.String())
	values = append(values, AssetID.HashID.String())
	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}

func (AssetID assetID) Equal(id types.ID) bool {
	switch id.(type) {
	case assetID:
		return bytes.Compare(AssetID.Bytes(), id.Bytes()) == 0
	default:
		return false
	}
}

func (assetID) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(assetID{}, module.Route+"/"+"assetID", nil)
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
func AssetIDAsKey(id types.ID) helpers.Key {
	return assetIDFromInterface(id)
}
func ReadClassificationID(assetID types.ID) types.ID {
	return assetIDFromInterface(assetID).ClassificationID
}
func ReadHashID(assetID types.ID) types.ID {
	return assetIDFromInterface(assetID).HashID
}

func NewAssetID(classificationID types.ID, immutables types.Immutables) types.ID {
	return assetID{
		ClassificationID: classificationID,
		HashID:           immutables.GetHashID(),
	}
}
