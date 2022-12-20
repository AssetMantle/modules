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

//type dataID struct {
//	Type ids.StringID
//	ids.HashId
//}

var _ ids.DataID = (*DataID)(nil)

func (dataID *DataID) GetHashID() ids.HashID {
	return dataID.HashId
}
func (dataID *DataID) IsDataID() {
}
func (dataID *DataID) DataIDString() string {
	return stringUtilities.JoinIDStrings(dataID.TypeId.String(), dataID.HashId.String())
}
func (dataID *DataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.TypeId.Bytes()...)
	Bytes = append(Bytes, dataID.HashId.Bytes()...)

	return Bytes
}
func (dataID *DataID) Compare(listable traits.Listable) int {
	return bytes.Compare(dataID.Bytes(), dataIDFromInterface(listable).Bytes())
}
func (dataID *DataID) ToAnyID() *AnyID {
	return &AnyID{
		Impl: &AnyID_DataId{
			DataId: dataID,
		},
	}
}

func dataIDFromInterface(i interface{}) *DataID {
	switch value := i.(type) {
	case *DataID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewDataID(data data.Data) ids.DataID {
	if data == nil {
		panic(errorConstants.MetaDataError)
	}

	return &DataID{
		TypeId: data.GetType().(*StringID),
		HashId: data.GenerateHashID().(*HashID),
	}
}

func PrototypeDataID() ids.DataID {
	return &DataID{
		TypeId: PrototypeStringID().(*StringID),
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadDataID(dataIDString string) (ids.DataID, error) {
	if typeAndHashIdString := stringUtilities.SplitCompositeIDString(dataIDString); len(typeAndHashIdString) == 2 {
		Type := NewStringID(typeAndHashIdString[0])
		if hashID, err := ReadHashID(typeAndHashIdString[1]); err == nil {
			return &DataID{
				TypeId: Type.(*StringID),
				HashId: hashID.(*HashID),
			}, nil
		}
	}

	if dataIDString == "" {
		return PrototypeDataID(), nil
	}

	return &DataID{}, errorConstants.MetaDataError
}
