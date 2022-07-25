// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type key struct {
	ids.ClassificationID
}

var _ helpers.Key = (*key)(nil)

func (key key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.Bytes())
}
func (key) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, key{})
}
func (key key) IsPartial() bool {
	return len(key.ClassificationID.Bytes()) == 0
}
func (key key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.Compare(CompareKey) == 0
	}
}
func keyFromInterface(i interface{}) (key, error) {
	switch value := i.(type) {
	case key:
		return value, nil
	default:
		return key{}, errorConstants.MetaDataError
	}
}

func FromID(id ids.ID) helpers.Key {
	if classificationID, err := keyFromInterface(id); err != nil {
		// TODO plug all panic scenarios
		panic(err)
	} else {
		return classificationID
	}
}

func Prototype() helpers.Key {
	return key{}
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

	return key{
		// TODO correct the initialization
		ClassificationID: baseIDs.NewStringID(stringUtilities.Hash(stringUtilities.Hash(immutableIDStringList...), stringUtilities.Hash(mutableIDStringList...), stringUtilities.Hash(defaultImmutableStringList...))),
	}
}
