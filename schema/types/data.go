/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Data interface {
	GetID() ID

	// Compare returns 1 if this > parameter
	// * returns -1 if this < parameter
	// * returns 0 if this = parameter
	Compare(Data) int

	String() string

	GetTypeID() ID
	ZeroValue() Data
	// GenerateHashID returns the hash of the Data as an ID
	// * Returns ID of empty string when the value of Data is that Data type's zero value
	GenerateHashID() ID

	AsAccAddress() (sdkTypes.AccAddress, error)
	AsListData() (ListData, error)
	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (ID, error)

	Get() interface{}

	Proto
}
