// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type classificationID struct {
	Hash ids.ID
}

var _ ids.ClassificationID = (*classificationID)(nil)
var _ helpers.Key = (*classificationID)(nil)

func (classificationID classificationID) Bytes() []byte {
	return classificationID.Hash.Bytes()
}
func (classificationID classificationID) String() string {
	return classificationID.Hash.String()
}
func (classificationID classificationID) Compare(listable traits.Listable) int {
	if compareClassificationID, err := classificationIDFromInterface(listable); err != nil {
		panic(classificationID)
	} else {
		return bytes.Compare(classificationID.Bytes(), compareClassificationID.Bytes())
	}
}
func (classificationID classificationID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(classificationID.Bytes())
}
func (classificationID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, classificationID{})
}
func (classificationID classificationID) IsPartial() bool {
	return len(classificationID.Hash.Bytes()) == 0
}
func (classificationID classificationID) Equals(key helpers.Key) bool {
	if compareClassificationID, err := classificationIDFromInterface(key); err != nil {
		return false
	} else {
		return classificationID.Compare(compareClassificationID) == 0
	}
}

func NewClassificationID(immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) ids.ClassificationID {
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
		if hashID := property.GetHash(); !(hashID.Compare(baseIDs.NewStringID("")) == 0) {
			defaultImmutableStringList[i] = hashID.String()
		}
	}

	return classificationID{
		Hash: baseIDs.NewStringID(stringUtilities.Hash(stringUtilities.Hash(immutableIDStringList...), stringUtilities.Hash(mutableIDStringList...), stringUtilities.Hash(defaultImmutableStringList...))),
	}
}
