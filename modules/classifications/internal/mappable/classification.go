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
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ mappables.Classification = (*Classification)(nil)

func (classification Classification) GetID() types.ID { return classification.ID }
func (classification Classification) GetKey() helpers.Key {
	return key.FromID(classification.ID)
}

func (classification Classification) GetImmutableProperties() types.Properties {
	return classification.HasImmutables.GetImmutableProperties()
}

func (classification Classification) GenerateHashID() types.ID {
	return classification.HasImmutables.GenerateHashID()
}

func (classification Classification) GetMutableProperties() types.Properties {
	return classification.HasMutables.GetMutableProperties()
}

func (classification Classification) Mutate(propertyList ...types.Property) traits.HasMutables {
	return classification.HasMutables.Mutate(propertyList...)
}

func (Classification) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, Classification{})
}

func NewClassification(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Classification {
	return &Classification{
		ID:            id,
		HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties},
		HasMutables:   baseTraits.HasMutables{Properties: mutableProperties},
	}
}
