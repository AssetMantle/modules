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
	return key.DataID.GetTypeID().Bytes()
}
func (key *Key) GenerateStoreKeyBytes() []byte {
	return key.DataID.GetHashID().Bytes()
}
func (key *Key) GeneratePrefixedStoreKeyBytes() []byte {
	return append(key.GenerateStorePrefixBytes(), key.GenerateStoreKeyBytes()...)
}
func (key *Key) IsPartial() bool {
	return len(key.DataID.GetHashID().Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.DataID.Compare(CompareKey.DataID) == 0
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

func NewKey(dataID ids.DataID) helpers.Key {
	return &Key{
		DataID: dataID.(*baseIDs.DataID),
	}
}

func Prototype() helpers.Key { return &Key{baseIDs.PrototypeDataID().(*baseIDs.DataID)} }
