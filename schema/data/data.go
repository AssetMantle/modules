// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO URI and PropertyID Data type
type Data interface {
	GetID() ids.DataID
	GetBondWeight() sdkTypes.Dec

	Unmarshal([]byte) error
	MarshalTo([]byte) (int, error)
	AsString() string
	FromString(string) (Data, error)
	Bytes() []byte

	GetType() ids.StringID
	ZeroValue() Data
	// GenerateHash returns the hash of the Data as an PropertyID
	// * Returns PropertyID of empty bytes when the value of Data is that Data type's zero value
	GenerateHashID() ids.HashID
	ToAnyData() AnyData
	traits.Listable
}
