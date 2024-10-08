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
	if key.SplitID != nil {
		if err := key.SplitID.ValidateBasic(); err != nil {
			return errorConstants.InvalidKey.Wrapf(err.Error())
		}
	}
	return nil
}
func (key *Key) GenerateStorePrefixBytes() []byte {
	return key.SplitID.GetAssetID().Bytes()
}
func (key *Key) GenerateStoreKeyBytes() []byte {
	return key.SplitID.GetOwnerID().Bytes()
}
func (key *Key) GeneratePrefixedStoreKeyBytes() []byte {
	return append(key.GenerateStorePrefixBytes(), key.GenerateStoreKeyBytes()...)
}
func (key *Key) IsPartial() bool {
	return len(key.SplitID.GetOwnerID().Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if CompareKey, ok := compareKey.(*Key); !ok {
		return false
	} else {
		return key.SplitID.Compare(CompareKey.SplitID) == 0
	}
}

func NewKey(splitID ids.SplitID) helpers.Key {
	return &Key{
		SplitID: splitID.(*baseIDs.SplitID),
	}
}

func Prototype() helpers.Key {
	return &Key{baseIDs.PrototypeSplitID().(*baseIDs.SplitID)}
}

func GetSplitIDFromKey(key helpers.Key) ids.SplitID {
	if key, ok := key.(*Key); !ok {
		return baseIDs.PrototypeSplitID().(*baseIDs.SplitID)
	} else {
		return key.SplitID
	}
}
