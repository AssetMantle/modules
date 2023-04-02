// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constants"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
	"strings"
)

var _ ids.PropertyID = (*PropertyID)(nil)

func (propertyID *PropertyID) ValidateBasic() error {
	if err := propertyID.KeyID.ValidateBasic(); err != nil {
		return err
	}
	if err := propertyID.TypeID.ValidateBasic(); err != nil {
		return err
	}
	return nil
}
func (propertyID *PropertyID) GetTypeID() ids.StringID {
	return NewStringID(constants.PropertyIDType)
}
func (propertyID *PropertyID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypePropertyID(), nil
	}

	keyIDAndTypeID := stringUtilities.SplitCompositeIDString(idString)
	if len(keyIDAndTypeID) != 2 {
		return PrototypePropertyID(), errorConstants.IncorrectFormat.Wrapf("expected composite id")
	} else if keyID, err := PrototypeStringID().FromString(keyIDAndTypeID[0]); err != nil {
		return PrototypePropertyID(), err
	} else if typeID, err := PrototypeStringID().FromString(keyIDAndTypeID[1]); err != nil {
		return PrototypePropertyID(), err
	} else {
		return &PropertyID{
			KeyID:  keyID.(*StringID),
			TypeID: typeID.(*StringID),
		}, nil
	}
}
func (propertyID *PropertyID) AsString() string {
	return stringUtilities.JoinIDStrings(propertyID.KeyID.AsString(), propertyID.TypeID.AsString())
}
func (propertyID *PropertyID) IsPropertyID() {}
func (propertyID *PropertyID) GetKey() ids.StringID {
	return propertyID.KeyID
}
func (propertyID *PropertyID) GetDataTypeID() ids.StringID {
	return propertyID.TypeID
}
func (propertyID *PropertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.KeyID.Bytes()...)
	Bytes = append(Bytes, propertyID.TypeID.Bytes()...)

	return Bytes
}
func (propertyID *PropertyID) Compare(listable traits.Listable) int {
	return bytes.Compare(propertyID.Bytes(), propertyIDFromInterface(listable).Bytes())
}
func (propertyID *PropertyID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_PropertyID{
			PropertyID: propertyID,
		},
	}
}

func propertyIDFromInterface(listable traits.Listable) *PropertyID {
	switch value := listable.(type) {
	case *PropertyID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("expected *PropertyID, got %T", listable))
	}
}
func PrototypePropertyID() ids.PropertyID {
	return &PropertyID{
		KeyID:  PrototypeStringID().(*StringID),
		TypeID: PrototypeStringID().(*StringID),
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return &PropertyID{
		KeyID:  key.(*StringID),
		TypeID: Type.(*StringID),
	}
}
