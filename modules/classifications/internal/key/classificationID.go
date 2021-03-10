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

type classificationID struct {
	ChainID types.ID `json:"chainID" valid:"required~required field chainID missing"`
	HashID  types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*classificationID)(nil)
var _ helpers.Key = (*classificationID)(nil)

func (classificationID classificationID) Bytes() []byte {
	return append(
		classificationID.ChainID.Bytes(),
		classificationID.HashID.Bytes()...)
}
func (classificationID classificationID) String() string {
	var values []string
	values = append(values, classificationID.ChainID.String())
	values = append(values, classificationID.HashID.String())

	return strings.Join(values, constants.IDSeparator)
}
func (classificationID classificationID) Equals(id types.ID) bool {
	return bytes.Equal(classificationID.Bytes(), id.Bytes())
}
func (classificationID classificationID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(classificationID.Bytes())
}
func (classificationID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, classificationID{})
}
func (classificationID classificationID) IsPartial() bool {
	return len(classificationID.HashID.Bytes()) == 0
}
func (classificationID classificationID) Matches(key helpers.Key) bool {
	return classificationID.Equals(classificationIDFromInterface(key))
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
		if hashID := property.GetFact().GetHashID(); !hashID.Equals(base.NewID("")) {
			defaultImmutableStringList[i] = hashID.String()
		}
	}

	return classificationID{
		ChainID: chainID,
		HashID:  base.NewID(metaUtilities.Hash(metaUtilities.Hash(immutableIDStringList...), metaUtilities.Hash(mutableIDStringList...), metaUtilities.Hash(defaultImmutableStringList...))),
	}
}
