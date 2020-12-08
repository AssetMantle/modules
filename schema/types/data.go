/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Data interface {
	String() string

	Type() string

	ZeroValue() Data

	GenerateHash() string

	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (ID, error)

	Get() interface{}

	Equal(Data) bool
}
