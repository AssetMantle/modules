// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ ids.AssetID = (*HashUser)(nil)

func (assetID HashUser) IsOwnableID() {}
func (assetID HashUser) IsAssetID()   {}

func assetIDFromInterface(i interface{}) *HashUser {
	switch value := i.(type) {
	case HashUser:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return &HashUser{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeAssetID() ids.AssetID {
	return &HashUser{
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return &HashUser{
			HashId: hashID.(*HashID),
		}, nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return &HashUser{}, errorConstants.MetaDataError
}
