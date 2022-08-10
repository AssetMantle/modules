// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

type dataID struct {
	Type ids.StringID
	ids.HashID
}

func (dataID dataID) IsDataID() {
	// TODO implement me
	panic("implement me")
}

var _ ids.DataID = (*dataID)(nil)

func (dataID dataID) String() string {
	return stringUtilities.JoinIDStrings(dataID.Type.String(), dataID.HashID.String())
}
func (dataID dataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.Type.Bytes()...)
	Bytes = append(Bytes, dataID.HashID.Bytes()...)

	return Bytes
}
func (dataID dataID) Compare(listable traits.Listable) int {
	return bytes.Compare(dataID.Bytes(), dataIDFromInterface(listable).Bytes())
}
func (dataID dataID) GetHashID() ids.HashID {
	return dataID.HashID
}
func dataIDFromInterface(i interface{}) dataID {
	switch value := i.(type) {
	case dataID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewDataID(data data.Data) ids.DataID {
	return dataID{
		Type:   data.GetType(),
		HashID: data.GenerateHashID(),
	}
}
