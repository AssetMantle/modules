// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/lists"
)

// TODO URI and ID data type
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
	AsDataList() (lists.DataList, error)
	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (ID, error)

	Get() interface{}
}
