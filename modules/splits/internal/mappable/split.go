/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type split struct {
	ID    types.ID     `json:"key" valid:"required field key missing"`
	Split sdkTypes.Dec `json:"split" valid:"required~required field split missing matches(^[0-9]$)~invalid field split"`
}

var _ mappables.Split = (*split)(nil)

func (split split) GetID() types.ID { return split.ID }
func (split split) GetOwnerID() types.ID {
	return key.ReadOwnerID(split.ID)
}
func (split split) GetOwnableID() types.ID {
	return key.ReadOwnableID(split.ID)
}
func (split split) GetSplit() sdkTypes.Dec {
	return split.Split
}
func (split split) Send(Split sdkTypes.Dec) traits.Transactional {
	split.Split = split.Split.Sub(Split)
	return split
}
func (split split) Receive(Split sdkTypes.Dec) traits.Transactional {
	split.Split = split.Split.Add(Split)
	return split
}
func (split split) CanSend(Split sdkTypes.Dec) bool {
	return split.Split.GTE(Split)
}
func (split split) GetKey() helpers.Key {
	return key.New(split.ID)
}
func (split) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(split{}, constants.ProjectRoute+"/"+"split", nil)
}

func NewSplit(splitID types.ID, spl sdkTypes.Dec) mappables.Split {
	return split{
		ID:    splitID,
		Split: spl,
	}
}
