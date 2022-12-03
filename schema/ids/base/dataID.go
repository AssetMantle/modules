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

var _ ids.DataID = (*DataID)(nil)

func (dataID DataID) IsDataID() {
}

// func (dataID DataID) String() string {
//	return stringUtilities.JoinIDStrings(dataID.Type.String(), dataID.HashID.String())
// }
func (dataID DataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.Type.Bytes()...)
	Bytes = append(Bytes, dataID.HashId.Bytes()...)

	return Bytes
}
func (dataID DataID) Compare(listable traits.Listable) int {
	return bytes.Compare(dataID.Bytes(), dataIDFromInterface(listable).Bytes())
}
func (dataID DataID) GetHashID() ids.HashID {
	return dataID.HashId
}
func dataIDFromInterface(i interface{}) DataID {
	switch value := i.(type) {
	case DataID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewDataID(data data.DataI) ids.DataID {
	if data == nil {
		panic(errorConstants.MetaDataError)
	}

	return &DataID{
		Type:   data.GetType().(*StringID),
		HashId: data.GenerateHashID().(*HashID),
	}
}

func PrototypeDataID() ids.DataID {
	return &DataID{
		Type:   PrototypeStringID().(*StringID),
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadDataID(dataIDString string) (ids.DataID, error) {
	if typeAndHashIDString := stringUtilities.SplitCompositeIDString(dataIDString); len(typeAndHashIDString) == 2 {
		Type := NewStringID(typeAndHashIDString[0])
		if hashID, err := ReadHashID(typeAndHashIDString[1]); err == nil {
			return &DataID{
				Type:   Type.(*StringID),
				HashId: hashID.(*HashID),
			}, nil
		}
	}

	if dataIDString == "" {
		return PrototypeDataID(), nil
	}

	return &DataID{}, errorConstants.MetaDataError
}
