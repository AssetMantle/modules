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
	if key.MaintainerID != nil {
		if err := key.MaintainerID.ValidateBasic(); err != nil {
			return errorConstants.InvalidKey.Wrapf(err.Error())
		}
	}
	return nil
}
func (key *Key) GenerateStorePrefixBytes() []byte {
	return []byte{0x0}
}
func (key *Key) GenerateStoreKeyBytes() []byte {
	return key.MaintainerID.Bytes()
}
func (key *Key) GeneratePrefixedStoreKeyBytes() []byte {
	return append(key.GenerateStorePrefixBytes(), key.GenerateStoreKeyBytes()...)
}
func (key *Key) IsPartial() bool {
	return len(key.MaintainerID.Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.MaintainerID.Compare(CompareKey.MaintainerID) == 0
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

func NewKey(maintainerID ids.MaintainerID) helpers.Key {
	return &Key{
		MaintainerID: maintainerID.(*baseIDs.MaintainerID),
	}
}

func Prototype() helpers.Key {
	return &Key{baseIDs.PrototypeMaintainerID().(*baseIDs.MaintainerID)}
}
