/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _ types.ID = (*ClassificationID)(nil)
var _ helpers.Key = (*ClassificationID)(nil)

func (classificationID ClassificationID) Bytes() []byte {
	return append(
		classificationID.ChainID.Bytes(),
		classificationID.HashID.Bytes()...)
}
func (classificationID ClassificationID) String() string {
	var values []string
	values = append(values, classificationID.ChainID.String())
	values = append(values, classificationID.HashID.String())

	return strings.Join(values, constants.IDSeparator)
}
func (classificationID ClassificationID) Compare(id types.ID) int {
	return bytes.Compare(classificationID.Bytes(), id.Bytes())
}
func (classificationID ClassificationID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(classificationID.Bytes())
}
func (ClassificationID) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, ClassificationID{})
}
func (classificationID ClassificationID) IsPartial() bool {
	return len(classificationID.HashID.Bytes()) == 0
}
func (classificationID ClassificationID) Equals(key helpers.Key) bool {
	return classificationID.Compare(classificationIDFromInterface(key)) == 0
}

func NewClassificationID(chainID types.ID, immutableProperties types.Properties, mutableProperties types.Properties) types.ID {
	immutableIDStringList := make([]string, len(immutableProperties.GetList()))

	for i, property := range immutableProperties.GetList() {
		immutableIDStringList[i] = property.GetID().String()
	}

	mutableIDStringList := make([]string, len(mutableProperties.GetList()))

	for i, property := range mutableProperties.GetList() {
		mutableIDStringList[i] = property.GetID().String()
	}

	defaultImmutableStringList := make([]string, len(immutableProperties.GetList()))

	for i, property := range immutableProperties.GetList() {
		if hashID := property.GetFact().GetHashID(); !(hashID.Compare(base.NewID("")) == 0) {
			defaultImmutableStringList[i] = hashID.String()
		}
	}

	return &ClassificationID{
		ChainID: chainID,
		HashID:  base.NewID(metaUtilities.Hash(metaUtilities.Hash(immutableIDStringList...), metaUtilities.Hash(mutableIDStringList...), metaUtilities.Hash(defaultImmutableStringList...))),
	}
}
