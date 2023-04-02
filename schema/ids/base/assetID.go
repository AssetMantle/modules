// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
	"strings"
)

var _ ids.AssetID = (*AssetID)(nil)

func (assetID *AssetID) ValidateBasic() error {
	return assetID.HashID.ValidateBasic()
}
func (assetID *AssetID) GetTypeID() ids.StringID {
	return NewStringID(constants.AssetIDType)
}
func (assetID *AssetID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeAssetID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeAssetID(), err
	} else {
		return &AssetID{
			HashID: hashID.(*HashID),
		}, nil
	}
}
func (assetID *AssetID) AsString() string {
	return assetID.HashID.AsString()
}
func (assetID *AssetID) Bytes() []byte {
	return assetID.HashID.IDBytes
}
func (assetID *AssetID) IsOwnableID() {}
func (assetID *AssetID) IsAssetID()   {}
func (assetID *AssetID) Compare(listable traits.Listable) int {
	compareID, err := idFromListable(listable)
	if err != nil {
		panic(err)
	}
	return bytes.Compare(assetID.Bytes(), compareID.Bytes())
}
func (assetID *AssetID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_AssetID{
			AssetID: assetID,
		},
	}
}
func (assetID *AssetID) ToAnyOwnableID() ids.AnyOwnableID {
	return &AnyOwnableID{
		Impl: &AnyOwnableID_AssetID{
			AssetID: assetID,
		},
	}
}

func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return &AssetID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeAssetID() ids.AssetID {
	return &AssetID{
		HashID: PrototypeHashID().(*HashID),
	}
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return &AssetID{
			HashID: hashID.(*HashID),
		}, nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return PrototypeAssetID(), errorConstants.MetaDataError.Wrapf("invalid assetID string: %s", assetIDString)
}
