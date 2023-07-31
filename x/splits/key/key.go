// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"

	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.Key = (*Key)(nil)

func (key *Key) GenerateStorePrefixBytes() []byte {
	return key.SplitID.GetOwnableID().Bytes()
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
	if compareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.SplitID.Compare(compareKey.SplitID) == 0
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

func NewKey(splitID ids.SplitID) helpers.Key {
	return &Key{
		SplitID: splitID.(*baseIDs.SplitID),
	}
}

func Prototype() helpers.Key {
	return &Key{baseIDs.PrototypeSplitID().(*baseIDs.SplitID)}
}

func GetSplitIDFromKey(key helpers.Key) ids.SplitID {
	if key, err := keyFromInterface(key); err != nil {
		return baseIDs.PrototypeSplitID().(*baseIDs.SplitID)
	} else {
		return key.SplitID
	}
}
