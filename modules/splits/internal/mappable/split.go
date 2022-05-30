// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type split struct {
	ID    types.ID     `json:"id" valid:"required field key missing"`
	Value sdkTypes.Dec `json:"value" valid:"required~required field value missing, matches(^[0-9]$)~invalid field value"`
}

var _ mappables.Split = (*split)(nil)

func (split split) GetID() types.ID { return split.ID }
func (split split) GetOwnerID() types.ID {
	return key.ReadOwnerID(split.ID)
}
func (split split) GetOwnableID() types.ID {
	return key.ReadOwnableID(split.ID)
}

// TODO rename to GetSupply
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
	return key.FromID(split.ID)
}
func (split) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, split{})
}

func NewSplit(splitID types.ID, value sdkTypes.Dec) mappables.Split {
	return split{
		ID:    splitID,
		Value: value,
	}
}
