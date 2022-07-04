// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type dataID struct {
	Type ids.ID
	Hash ids.ID
}

var _ ids.DataID = (*dataID)(nil)

func (dataID dataID) String() string {
	return stringUtilities.JoinIDStrings(dataID.Type.String(), dataID.Hash.String())
}
func (dataID dataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.Type.Bytes()...)
	Bytes = append(Bytes, dataID.Hash.Bytes()...)

	return Bytes
}
func (dataID dataID) Compare(listable traits.Listable) int {
	if compareDataID, err := dataIDFromInterface(listable); err != nil {
		panic(errorConstants.MetaDataError)
	} else {
		return bytes.Compare(dataID.Bytes(), compareDataID.Bytes())
	}
}
func (dataID dataID) GetHash() ids.ID {
	return dataID.Hash
}
func dataIDFromInterface(i interface{}) (dataID, error) {
	switch value := i.(type) {
	case dataID:
		return value, nil
	default:
		return dataID{}, errorConstants.MetaDataError
	}
}
func NewDataID(data data.Data) ids.DataID {
	return dataID{
		Type: data.GetType(),
		Hash: data.GenerateHash(),
	}
}
