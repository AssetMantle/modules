// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ ids.AssetID = (*HashID)(nil)

func (assetID HashID) IsOwnableID() {}
func (assetID HashID) IsAssetID()   {}

func assetIDFromInterface(i interface{}) *HashID {
	switch value := i.(type) {
	case HashID:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return &HashID{
		HashBytes: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).Bytes(),
	}
}

func PrototypeAssetID() ids.AssetID {
	return &HashID{
		HashBytes: PrototypeHashID().Bytes(),
	}
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return &HashID{
			HashBytes: hashID.Bytes(),
		}, nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return &HashID{}, errorConstants.MetaDataError
}
