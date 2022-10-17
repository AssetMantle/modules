// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type mappable struct {
	types.Identity
}

var _ types.Identity = (*mappable)(nil)

func (identity mappable) GetKey() helpers.Key {
	return key.NewKey(baseIDs.NewIdentityID(identity.GetClassificationID(), identity.GetImmutables()))
}
func (mappable) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, mappable{})
}

func NewMappable(identity types.Identity) helpers.Mappable {
	return mappable{
		Identity: identity,
	}
}

func Prototype() helpers.Mappable {
	return mappable{}
}
