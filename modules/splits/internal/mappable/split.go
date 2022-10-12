// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type split struct {
	OwnerID   ids.IdentityID
	OwnableID ids.OwnableID
	Value     sdkTypes.Dec
}

var _ mappables.Split = (*split)(nil)

func (split split) GetOwnerID() ids.ID {
	return split.OwnerID
}
func (split split) GetOwnableID() ids.ID {
	return split.OwnableID
}
func (split split) GetValue() sdkTypes.Dec {
	return split.Value
}
func (split split) Send(outValue sdkTypes.Dec) capabilities.Transactional {
	split.Value = split.Value.Sub(outValue)
	return split
}
func (split split) Receive(inValue sdkTypes.Dec) capabilities.Transactional {
	split.Value = split.Value.Add(inValue)
	return split
}
func (split split) CanSend(outValue sdkTypes.Dec) bool {
	return split.Value.GTE(outValue)
}
func (split split) GetKey() helpers.Key {
	return key.NewKey(base.NewSplitID(split.OwnerID, split.OwnableID))
}
func (split) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, split{})
}

func NewSplit(ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) mappables.Split {
	return split{
		OwnerID:   ownerID,
		OwnableID: ownableID,
		Value:     value,
	}
}

func Prototype() helpers.Mappable {
	return split{}
}
