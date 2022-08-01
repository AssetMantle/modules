// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type asset struct {
	qualified.Document
}

var _ mappables.Asset = (*asset)(nil)

func (asset asset) GetBurn() properties.Property {
	if burn := asset.GetProperty(constants.BurnProperty); burn != nil {
		return burn
	}

	return constants.Burn
}
func (asset asset) GetLock() properties.Property {
	if lock := asset.GetProperty(constants.LockProperty); lock != nil {
		return lock
	}

	return constants.Lock
}
func (asset asset) GetSupply() properties.Property {
	if supply := asset.GetProperty(constants.SupplyProperty); supply != nil {
		return supply
	}

	return constants.Supply
}
func (asset asset) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewAssetID(asset.GetClassificationID(), asset.GetImmutables()))
}
func (asset) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, asset{})
}

func NewAsset(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) mappables.Asset {
	return asset{
		Document: baseQualified.NewDocument(classificationID, immutables, mutables),
	}
}

func Prototype() helpers.Mappable {
	return asset{}
}
