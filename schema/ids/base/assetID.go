// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

// type assetID struct {
//	ids.HashID
// }

var _ ids.AssetID = (*AssetID)(nil)

func (assetID *AssetID) Bytes() []byte {
	return assetID.HashID.IDBytes
}
func (assetID *AssetID) IsOwnableID() {}
func (assetID *AssetID) IsAssetID()   {}
func (assetID *AssetID) Compare(listable traits.Listable) int {
	// TODO devise a better strategy to compare assetID and ownableID
	return bytes.Compare(assetID.Bytes(), ownableIDFromInterface(listable).Bytes())
}
func (assetID *AssetID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_AssetID{
			AssetID: assetID,
		},
	}
}
func assetIDFromInterface(i interface{}) *AssetID {
	switch value := i.(type) {
	case *AssetID:
		return value
	default:
		panic(errorConstants.MetaDataError)
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

	return &AssetID{}, errorConstants.MetaDataError
}
