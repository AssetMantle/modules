// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	metaUtilities "github.com/AssetMantle/modules/utilities/meta"
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
func (classificationID classificationID) Compare(id types.ID) int {
	return bytes.Compare(classificationID.Bytes(), id.Bytes())
}
func (classificationID classificationID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(classificationID.Bytes())
}
func (classificationID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, classificationID{})
}
func (classificationID classificationID) IsPartial() bool {
	return len(classificationID.HashID.Bytes()) == 0
}
func (classificationID classificationID) Equals(key helpers.Key) bool {
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
		if hashID := property.GetHash(); !(hashID.Compare(baseIDs.NewID("")) == 0) {
			defaultImmutableStringList[i] = hashID.String()
		}
	}

	return classificationID{
		ChainID: chainID,
		HashID:  baseIDs.NewID(metaUtilities.Hash(metaUtilities.Hash(immutableIDStringList...), metaUtilities.Hash(mutableIDStringList...), metaUtilities.Hash(defaultImmutableStringList...))),
	}
}
