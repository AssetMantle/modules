// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type mappable struct {
	documents.Maintainer
}

var _ helpers.Mappable = (*mappable)(nil)

func (maintainer mappable) GetKey() helpers.Key {
	return key.NewKey(base.NewMaintainerID(maintainer.GetMaintainedClassificationID(), maintainer.GetIdentityID()))
}
func (mappable) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, mappable{})
}

func NewMappable(maintainer documents.Maintainer) helpers.Mappable {
	return mappable{Maintainer: maintainer}
}

func Prototype() helpers.Mappable {
	return mappable{}
}
