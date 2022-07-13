// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type assetID struct {
	ids.ClassificationID
	Hash ids.ID
}

var _ ids.AssetID = (*assetID)(nil)
var _ helpers.Key = (*assetID)(nil)

func (assetID assetID) String() string {
	return stringUtilities.JoinIDStrings(assetID.ClassificationID.String(), assetID.Hash.String())
}
func (assetID assetID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, assetID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, assetID.Hash.Bytes()...)

	return Bytes
}
func (assetID assetID) Compare(listable traits.Listable) int {
	if compareAssetID, err := assetIDFromInterface(listable); err != nil {
		panic(err)
	} else {
		return bytes.Compare(assetID.Bytes(), compareAssetID.Bytes())
	}
}
func (assetID assetID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(assetID.Bytes())
}
func (assetID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, assetID{})
}
func (assetID assetID) IsPartial() bool {
	return len(assetID.Hash.Bytes()) == 0
}
func (assetID assetID) Equals(key helpers.Key) bool {
	if compareAssetID, err := assetIDFromInterface(key); err != nil {
		return false
	} else {
		return assetID.Compare(compareAssetID) == 0
	}
}

func NewAssetID(classificationID ids.ClassificationID, immutableProperties lists.PropertyList) ids.AssetID {
	return assetID{
		ClassificationID: classificationID,
		Hash:             base.Immutables{PropertyList: immutableProperties}.GenerateHashID(),
	}
}
