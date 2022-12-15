// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type mappable struct {
	documents.Maintainer
}

var _ helpers.Mappable = (*mappable)(nil)

func (mappable mappable) GetKey() helpers.Key {
	return key.NewKey(base.NewMaintainerID(constansts.MaintainerClassificationID,
		baseQualified.NewImmutables(baseLists.NewPropertyList(
			baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(mappable.Maintainer.GetMaintainedClassificationID())),
			baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(mappable.Maintainer.GetIdentityID())),
		))))
}
func (mappable) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, mappable{})
}

func NewMappable(maintainer documents.Maintainer) helpers.Mappable {
	return mappable{Maintainer: maintainer}
}

func Prototype() helpers.Mappable {
	return mappable{}
}
