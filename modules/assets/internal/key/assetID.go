/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/traits/base"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ types.ID = (*AssetID)(nil)
var _ helpers.Key = (*AssetID)(nil)

func (assetID AssetID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, assetID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, assetID.HashID.Bytes()...)

	return Bytes
}
func (assetID AssetID) String() string {
	var values []string
	values = append(values, assetID.ClassificationID.String())
	values = append(values, assetID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (assetID AssetID) Compare(id types.ID) int {
	return bytes.Compare(assetID.Bytes(), id.Bytes())
}
func (assetID AssetID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(assetID.Bytes())
}
func (AssetID) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, AssetID{})
}
func (assetID AssetID) IsPartial() bool {
	return len(assetID.HashID.Bytes()) == 0
}
func (assetID AssetID) Equals(key helpers.Key) bool {
	id := assetIDFromInterface(key)
	return assetID.Compare(&id) == 0
}

// NewAssetID TODO try removing & like in mappables
func NewAssetID(classificationID types.ID, immutableProperties types.Properties) types.ID {
	return &AssetID{
		ClassificationID: classificationID,
		HashID:           base.HasImmutables{Properties: immutableProperties}.GenerateHashID(),
	}
}
