// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type assetID base.AssetID

var _ ids.AssetID = (*assetID)(nil)

func (assetID *assetID) String() string {
	return assetID.String()
}
func (assetID *assetID) Compare(listable traits.Listable) int {
	return bytes.Compare(assetID.Bytes(), assetIDFromInterface(listable).HashId.Bytes())

}
func (assetID *assetID) IsOwnableID() {}
func (assetID *assetID) IsAssetID()   {}

func assetIDFromInterface(i interface{}) *assetID {
	switch value := i.(type) {
	case assetID:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func (assetID *assetID) Bytes() []byte {
	return assetID.HashId.GetIdBytes()
}
func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return &assetID{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()),
	}
}

func PrototypeAssetID() ids.AssetID {
	return &assetID{
		HashId: PrototypeHashID(),
	}
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return &assetID{
			HashId: hashID,
		}, nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return &assetID{}, errorConstants.MetaDataError
}
