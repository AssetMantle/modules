// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids/constants"
	"strings"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.StringID = (*StringID)(nil)

func (stringID *StringID) ValidateBasic() error {
	if !utilities.IsValidStringID(stringID.AsString()) {
		return errorConstants.IncorrectFormat
	}
	return nil
}
func (stringID *StringID) IsStringID() {}
func (stringID *StringID) GetTypeID() ids.StringID {
	return NewStringID(constants.StringIDType)
}
func (stringID *StringID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeStringID(), nil
	}

	returnStringID := &StringID{
		IDString: idString,
	}

	if err := returnStringID.ValidateBasic(); err != nil {
		return PrototypeStringID(), err
	} else {
		return returnStringID, nil
	}
}
func (stringID *StringID) AsString() string {
	return stringID.IDString
}
func (stringID *StringID) Bytes() []byte {
	return []byte(stringID.IDString)
}
func (stringID *StringID) Compare(listable traits.Listable) int {
	return strings.Compare(stringID.AsString(), stringIDFromInterface(listable).AsString())
}
func (stringID *StringID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_StringID{
			StringID: stringID,
		},
	}
}

func stringIDFromInterface(i interface{}) *StringID {
	switch value := i.(type) {
	case *StringID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("expected *StringID, got %T", i))
	}
}

func NewStringID(idString string) ids.StringID {
	return &StringID{IDString: idString}
}

func PrototypeStringID() ids.StringID {
	return &StringID{
		IDString: "",
	}
}
