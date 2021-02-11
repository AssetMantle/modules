/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Data interface {
	String() string

	GetTypeID() ID

	ZeroValue() Data

	GenerateHashID() ID

	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (ID, error)

	Get() interface{}

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(bz []byte) error

	MarshalAmino() (string, error)
	UnmarshalAmino(text string) (err error)

	Equal(Data) bool
}
