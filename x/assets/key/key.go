// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"

	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.Key = (*Key)(nil)

func (key *Key) GenerateStorePrefixBytes() []byte {
	return []byte{}
}
func (key *Key) GenerateStoreKeyBytes() []byte {
	return key.AssetID.Bytes()
}
func (key *Key) GeneratePrefixedStoreKeyBytes() []byte {
	return append(key.GenerateStorePrefixBytes(), key.GenerateStoreKeyBytes()...)
}
func (key *Key) IsPartial() bool {
	return len(key.AssetID.Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		// TODO test nil AssetID case
		return key.AssetID.Compare(CompareKey.AssetID) == 0
	}
}
func keyFromInterface(i interface{}) (*Key, error) {
	switch value := i.(type) {
	case *Key:
		return value, nil
	default:
		return &Key{}, errorConstants.IncorrectFormat.Wrapf("incorrect key type")
	}
}

func NewKey(assetID ids.AssetID) helpers.Key {
	return &Key{
		AssetID: assetID.(*baseIDs.AssetID),
	}
}

func Prototype() helpers.Key {
	return &Key{baseIDs.PrototypeAssetID().(*baseIDs.AssetID)}
}
