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

//
// type dataID struct {
//	Type ids.StringID
//	ids.HashID
// }

var _ ids.DataID = (*DataIDI_DataID)(nil)

func (dataID *DataIDI_DataID) IsDataID() {
}
func (dataID *DataIDI_DataID) String() string {
	return stringUtilities.JoinIDStrings(dataID.DataID.Type.String(), dataID.DataID.HashId.String())
}

// func (dataID *dataID) String() string {
//	return stringUtilities.JoinIDStrings(dataID.Type.String(), dataID.HashID.String())
// }
func (dataID *DataIDI_DataID) Bytes() []byte {
	return append(dataID.DataID.Type.Bytes(), dataID.DataID.HashId.Bytes()...)
}
func (dataID *DataIDI_DataID) Compare(listable traits.Listable) int {
	return bytes.Compare(dataID.Bytes(), dataIDFromInterface(listable).Bytes())
}
func (dataID *DataIDI_DataID) GetHashID() ids.HashID {
	return dataID.DataID.HashId
}
func dataIDFromInterface(i interface{}) *DataIDI_DataID {
	switch value := i.(type) {
	case *DataIDI_DataID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func GenerateDataID(data data.Data) ids.DataID {
	if data == nil {
		panic(errorConstants.MetaDataError)
	}
	return NewDataID(data.GetType(), data.GenerateHashID())
}
func NewDataID(Type ids.StringID, hashID ids.HashID) ids.DataID {
	return &DataIDI{
		Impl: &DataIDI_DataID{
			DataID: &DataID{
				Type:   Type.(*StringIDI),
				HashId: hashID.(*HashIDI),
			},
		},
	}
}

func PrototypeDataID() ids.DataID {
	return NewDataID(PrototypeStringID(), PrototypeHashID())
}

func ReadDataID(dataIDString string) (ids.DataID, error) {
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
