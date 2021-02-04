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
	ID                  types.ID         `json:"id" valid:"required~required field id missing"`
	ImmutableProperties types.Properties `json:"immutableProperties" valid:"required field immutableProperties missing"`
	MutableProperties   types.Properties `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
}

var _ mappables.Classification = (*classification)(nil)

func (classification classification) GetID() types.ID { return classification.ID }
func (classification classification) GetImmutableProperties() types.Properties {
	return classification.ImmutableProperties
}
func (classification classification) GetMutableProperties() types.Properties {
	return classification.MutableProperties
}
func (classification classification) GetKey() helpers.Key {
	return key.New(classification.ID)
}

func (classification) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, classification{})
}

func NewClassification(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Classification {
	return classification{
		ID:                  id,
		ImmutableProperties: immutableProperties,
		MutableProperties:   mutableProperties,
	}
}
