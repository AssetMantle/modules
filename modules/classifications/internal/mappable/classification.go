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
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type classification struct {
	ID types.ID `json:"id" valid:"required~required field id missing"`
	baseTraits.HasImmutables
	baseTraits.HasMutables //nolint:govet
}

var _ mappables.Classification = (*classification)(nil)

func (classification classification) GetID() types.ID { return classification.ID }
func (classification classification) GetKey() helpers.Key {
	return key.FromID(classification.ID)
}

func (classification) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, classification{})
}

func NewClassification(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Classification {
	return classification{
		ID:            id,
		HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
		HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
	}
}
