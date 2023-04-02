// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constants"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
	"strings"
)

var _ ids.DataID = (*DataID)(nil)

func (dataID *DataID) ValidateBasic() error {
	if err := dataID.TypeID.ValidateBasic(); err != nil {
		return err
	}
	if err := dataID.HashID.ValidateBasic(); err != nil {
		return err
	}
	return nil
}
func (dataID *DataID) GetTypeID() ids.StringID {
	return NewStringID(constants.DataIDType)
}
func (dataID *DataID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeDataID(), nil
	}

	typeIDAndHashID := stringUtilities.SplitCompositeIDString(idString)
	if len(typeIDAndHashID) != 2 {
		return PrototypeDataID(), errorConstants.IncorrectFormat.Wrapf("expected composite id")
	} else if typeID, err := PrototypeStringID().FromString(typeIDAndHashID[0]); err != nil {
		return PrototypeDataID(), err

	} else if hashID, err := PrototypeHashID().FromString(typeIDAndHashID[1]); err != nil {
		return PrototypeDataID(), err
	} else {
		return &DataID{
			TypeID: typeID.(*StringID),
			HashID: hashID.(*HashID),
		}, nil
	}
}
func (dataID *DataID) AsString() string {
	return stringUtilities.JoinIDStrings(dataID.TypeID.AsString(), dataID.HashID.AsString())
}
func (dataID *DataID) GetHashID() ids.HashID {
	return dataID.HashID
}
func (dataID *DataID) IsDataID() {
}
func (dataID *DataID) DataIDString() string {
	return stringUtilities.JoinIDStrings(dataID.TypeID.AsString(), dataID.HashID.AsString())
}
func (dataID *DataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.TypeID.Bytes()...)
	Bytes = append(Bytes, dataID.HashID.Bytes()...)

	return Bytes
}
func (dataID *DataID) Compare(listable traits.Listable) int {
	return bytes.Compare(dataID.Bytes(), dataIDFromInterface(listable).Bytes())
}
func (dataID *DataID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_DataID{
			DataID: dataID,
		},
	}
}

func dataIDFromInterface(i interface{}) *DataID {
	switch value := i.(type) {
	case *DataID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("expected *DataID, got %T", i))
	}
}

func GenerateDataID(data data.Data) ids.DataID {
	if data == nil {
		panic(errorConstants.MetaDataError.Wrapf("data is nil"))
	}

	return &DataID{
		TypeID: data.GetTypeID().(*StringID),
		HashID: data.GenerateHashID().(*HashID),
	}
}

func PrototypeDataID() ids.DataID {
	return &DataID{
		TypeID: PrototypeStringID().(*StringID),
		HashID: PrototypeHashID().(*HashID),
	}
}

func ReadDataID(dataIDString string) (ids.DataID, error) {
	if typeAndHashIdString := stringUtilities.SplitCompositeIDString(dataIDString); len(typeAndHashIdString) == 2 {
		Type := NewStringID(typeAndHashIdString[0])
		if hashID, err := ReadHashID(typeAndHashIdString[1]); err == nil {
			return &DataID{
				TypeID: Type.(*StringID),
				HashID: hashID.(*HashID),
			}, nil
		}
	}

	if dataIDString == "" {
		return PrototypeDataID(), nil
	}

	return PrototypeDataID(), errorConstants.MetaDataError.Wrapf("invalid dataIDString: %s", dataIDString)
}
