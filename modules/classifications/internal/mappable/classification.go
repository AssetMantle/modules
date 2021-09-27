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
	"github.com/persistenceOne/persistenceSDK/schema/mappables" //nolint:typecheck
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	baseTypes "github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ mappables.Classification = (*Classification)(nil)

func (classification Classification) GetStructReference() codec.ProtoMarshaler {
	return &classification
}
func (classification Classification) GetID() types.ID { return &classification.ID }
func (classification Classification) GetKey() helpers.Key {
	return key.FromID(&classification.ID)
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
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, Classification{})
}

func NewClassification(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Classification {
	return &Classification{
		ID:            *baseTypes.NewID(id.String()),
		HasImmutables: baseTraits.HasImmutables{Properties: *baseTypes.NewProperties(immutableProperties.GetList()...)},
		HasMutables:   baseTraits.HasMutables{Properties: *baseTypes.NewProperties(mutableProperties.GetList()...)},
	}
}
