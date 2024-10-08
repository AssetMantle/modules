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
	if key.OrderID != nil {
		if err := key.OrderID.ValidateBasic(); err != nil {
			return errorConstants.InvalidKey.Wrapf(err.Error())
		}
	}
	return nil
}
func (key *Key) GenerateStorePrefixBytes() []byte {
	return []byte{0x0}
}
func (key *Key) GenerateStoreKeyBytes() []byte {
	return key.OrderID.Bytes()
}
func (key *Key) GeneratePrefixedStoreKeyBytes() []byte {
	return append(key.GenerateStorePrefixBytes(), key.GenerateStoreKeyBytes()...)
}
func (key *Key) IsPartial() bool {
	return len(key.OrderID.Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if CompareKey, ok := compareKey.(*Key); !ok {
		return false
	} else {
		return key.OrderID.Compare(compareKey.OrderID) == 0
	}
}

func NewKey(orderID ids.OrderID) helpers.Key {
	return &Key{
		OrderID: orderID.(*baseIDs.OrderID),
	}
}

func Prototype() helpers.Key {
	return &Key{baseIDs.PrototypeOrderID().(*baseIDs.OrderID)}
}
