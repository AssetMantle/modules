/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
	"strings"
)

type classificationID struct {
	ChainID types.ID `json:"chainID" valid:"required~required field chainID missing"`
	HashID  types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*classificationID)(nil)

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
func ChainIDFromClassificationID(classificationID types.ID) types.ID {
	return classificationIDFromInterface(classificationID).ChainID
}
func readClassificationID(classificationIDString string) types.ID {
	idList := strings.Split(classificationIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return classificationID{
			ChainID: base.NewID(idList[0]),
			HashID:  base.NewID(idList[1]),
		}
	}
	return classificationID{ChainID: base.NewID(""), HashID: base.NewID("")}
}
func classificationIDFromInterface(id types.ID) classificationID {
	switch value := id.(type) {
	case classificationID:
		return value
	default:
		return classificationIDFromInterface(readClassificationID(id.String()))
	}
}
func generateKey(classificationID types.ID) []byte {
	return append(StoreKeyPrefix, classificationIDFromInterface(classificationID).Bytes()...)
}
func NewClassificationID(chainID types.ID, immutableTraits types.Immutables, mutableTraits types.Mutables) types.ID {
	var immutableIDStringList []string
	for _, immutable := range immutableTraits.Get().GetList() {
		immutableIDStringList = append(immutableIDStringList, immutable.GetID().String())
	}
	var mutableIDStringList []string
	for _, mutable := range mutableTraits.Get().GetList() {
		mutableIDStringList = append(immutableIDStringList, mutable.GetID().String())
	}
	return classificationID{
		ChainID: chainID,
		HashID:  base.NewID(metaUtilities.Hash(metaUtilities.Hash(immutableIDStringList...), metaUtilities.Hash(mutableIDStringList...), immutableTraits.GetHashID().String())),
	}
}
