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

var _ ids.AssetID = (*ID_AssetID)(nil)

func (assetID *ID_AssetID) String() string {
	return assetID.AssetID.String()
}
func (assetID *ID_AssetID) Compare(listable traits.Listable) int {
	return bytes.Compare(assetID.Bytes(), idFromInterface(listable).Bytes())

}
func (assetID *ID_AssetID) IsOwnableID() {}
func (assetID *ID_AssetID) IsAssetID()   {}
func (assetID *ID_AssetID) Bytes() []byte {
	return assetID.AssetID.HashId.IdBytes
}
func GenerateAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.ID {
	return NewAssetID(GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()))
}

func NewAssetID(hashID ids.ID) ids.ID {
	if hashID.(*ID).GetHashID() == nil {
		panic(errorConstants.MetaDataError)
	}

	return &ID{
		Impl: &ID_AssetID{
			AssetID: &AssetID{
				HashId: hashID.(*ID).GetHashID(),
			},
		},
	}
}

func PrototypeAssetID() ids.ID {
	return NewAssetID(PrototypeHashID())
}

func ReadAssetID(assetIDString string) (ids.ID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return NewAssetID(hashID), nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return PrototypeAssetID(), errorConstants.MetaDataError
}
