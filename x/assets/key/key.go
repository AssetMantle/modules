// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var _ helpers.Key = (*Key)(nil)

func (key *Key) ValidateBasic() error {
	if key.AssetID != nil {
		if err := key.AssetID.ValidateBasic(); err != nil {
			return errorConstants.InvalidKey.Wrapf(err.Error())
		}
	}
	return nil
}
func (key *Key) GenerateStorePrefixBytes() []byte {
	return []byte{0x0}
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
	if CompareKey, ok := compareKey.(*Key); !ok {
		return false
	} else {
		// TODO test nil AssetID case
		return key.AssetID.Compare(CompareKey.AssetID) == 0
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
