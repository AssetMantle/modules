/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
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

func (classification classification) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(classification)
}
func (classification classification) Decode(bytes []byte) helpers.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &classification)
	return classification
}

func classificationPrototype() helpers.Mappable {
	return classification{}
}

func NewClassification(ID types.ID, immutableTraits types.Immutables, mutableTraits types.Mutables) mappables.Classification {
	return classification{
		ID:              ID,
		ImmutableTraits: immutableTraits,
		MutableTraits:   mutableTraits,
	}
}
