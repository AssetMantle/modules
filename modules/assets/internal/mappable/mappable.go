// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type mappable struct {
	types.Asset
}

var _ helpers.Mappable = (*mappable)(nil)

func (asset mappable) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewAssetID(asset.GetClassificationID(), asset.GetImmutables()))
}
func (mappable) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, mappable{})
}

func Prototype() helpers.Mappable {
	return mappable{}
}

func NewMappable(asset types.Asset) helpers.Mappable {
	return mappable{
		Asset: asset,
	}
}
