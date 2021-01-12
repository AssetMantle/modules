/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type classification struct {
	ID              types.ID         `json:"id" valid:"required~required field id missing"`
	ImmutableTraits types.Immutables `json:"immutableTraits" valid:"required field immutableTraits missing"`
	MutableTraits   types.Mutables   `json:"mutableTraits" valid:"required~required field mutableTraits missing"`
}

var _ mappables.Classification = (*classification)(nil)

func (classification classification) GetID() types.ID { return classification.ID }
func (classification classification) GetImmutables() types.Immutables {
	return classification.ImmutableTraits
}
func (classification classification) GetMutables() types.Mutables {
	return classification.MutableTraits
}
func (classification classification) GetKey() helpers.Key {
	return key.New(classification.ID)
}

func (classification) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, classification{})
}

func NewClassification(ID types.ID, immutableTraits types.Immutables, mutableTraits types.Mutables) mappables.Classification {
	return classification{
		ID:              ID,
		ImmutableTraits: immutableTraits,
		MutableTraits:   mutableTraits,
	}
}
