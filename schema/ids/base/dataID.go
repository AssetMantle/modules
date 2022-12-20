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

var _ ids.DataID = (*ID_DataID)(nil)

func (dataID *ID_DataID) IsDataID() {
}
func (dataID *ID_DataID) String() string {
	return stringUtilities.JoinIDStrings(dataID.DataID.Type.String(), dataID.DataID.HashID.String())
}
func (dataID *ID_DataID) Bytes() []byte {
	return append([]byte(dataID.DataID.Type.IdString), dataID.DataID.HashID.IdBytes...)
}
func (dataID *ID_DataID) Compare(listable traits.Listable) int {
	return bytes.Compare(dataID.Bytes(), idFromInterface(listable).Bytes())
}
func (dataID *ID_DataID) GetHashID() ids.ID {
	return &ID{Impl: &ID_HashID{HashID: dataID.DataID.HashID}}
}

func GenerateDataID(data data.Data) ids.ID {
	if data == nil {
		panic(errorConstants.MetaDataError)
	}
	return NewDataID(data.GetType(), data.GenerateHashID())
}
func NewDataID(Type ids.ID, hashID ids.ID) ids.ID {
	if Type.(*ID).GetStringID() == nil || hashID.(*ID).GetHashID() == nil {
		panic(errorConstants.MetaDataError)
	}
	return &ID{
		Impl: &ID_DataID{
			DataID: &DataID{
				Type:   Type.(*ID).GetStringID(),
				HashID: hashID.(*ID).GetHashID(),
			},
		},
	}
}

func PrototypeDataID() ids.ID {
	return NewDataID(PrototypeStringID(), PrototypeHashID())
}

func ReadDataID(dataIDString string) (ids.ID, error) {
	if typeAndHashIDString := stringUtilities.SplitCompositeIDString(dataIDString); len(typeAndHashIDString) == 2 {
		Type := NewStringID(typeAndHashIDString[0])
		if hashID, err := ReadHashID(typeAndHashIDString[1]); err == nil {
			return NewDataID(Type, hashID), nil
		}
	}

	if dataIDString == "" {
		return PrototypeDataID(), nil
	}

	return PrototypeDataID(), errorConstants.MetaDataError
}
