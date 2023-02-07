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

type assetID struct {
	ids.HashID
}

var _ ids.AssetID = (*assetID)(nil)

func (assetID assetID) IsOwnableID() {}
func (assetID assetID) IsAssetID()   {}
func (assetID assetID) Compare(listable traits.Listable) int {
	// TODO devise a better strategy to compare assetID and ownableID
	return bytes.Compare(assetID.Bytes(), ownableIDFromInterface(listable).Bytes())
}
func assetIDFromInterface(i interface{}) assetID {
	switch value := i.(type) {
	case assetID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return assetID{
		GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()),
	}
}

func PrototypeAssetID() ids.AssetID {
	return assetID{
		HashID: PrototypeHashID(),
	}
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return assetID{
			HashID: hashID,
		}, nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return assetID{}, errorConstants.MetaDataError
}
