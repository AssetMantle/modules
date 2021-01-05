/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
	"strings"
)

type classificationID struct {
	ChainID types.ID `json:"chainID" valid:"required~required field chainID missing"`
	HashID  types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*classificationID)(nil)
var _ helpers.Key = (*classificationID)(nil)

func (ClassificationID classificationID) Bytes() []byte {
	return append(
		ClassificationID.ChainID.Bytes(),
		ClassificationID.HashID.Bytes()...)
}
func (ClassificationID classificationID) String() string {
	var values []string
	values = append(values, ClassificationID.ChainID.String())
	values = append(values, ClassificationID.HashID.String())
	return strings.Join(values, constants.IDSeparator)
}
func (ClassificationID classificationID) Equals(id types.ID) bool {
	return bytes.Compare(ClassificationID.Bytes(), id.Bytes()) == 0
}
func (ClassificationID classificationID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(ClassificationID.Bytes())
}
func (classificationID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, classificationID{})
}
func (ClassificationID classificationID) IsPartial() bool {
	if len(ClassificationID.HashID.Bytes()) > 0 {
		return false
	}
	return true
}
func (ClassificationID classificationID) Matches(key helpers.Key) bool {
	switch value := key.(type) {
	case classificationID:
		return bytes.Compare(ClassificationID.Bytes(), value.Bytes()) == 0
	default:
		return false
	}
}

func New(id types.ID) helpers.Key {
	return classificationIDFromInterface(id)
}

func NewClassificationID(chainID types.ID, immutableTraits types.Immutables, mutableTraits types.Mutables) types.ID {
	var immutableIDStringList []string
	for _, immutable := range immutableTraits.Get().GetList() {
		immutableIDStringList = append(immutableIDStringList, immutable.GetID().String())
	}
	var mutableIDStringList []string
	for _, mutable := range mutableTraits.Get().GetList() {
		mutableIDStringList = append(mutableIDStringList, mutable.GetID().String())
	}
	return classificationID{
		ChainID: chainID,
		HashID:  base.NewID(metaUtilities.Hash(metaUtilities.Hash(immutableIDStringList...), metaUtilities.Hash(mutableIDStringList...), immutableTraits.GetHashID().String())),
	}
}
