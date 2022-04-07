// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/modules/schema/ids"
)

// TODO URI and ID data type
type Data interface {
	GetID() ids.DataID

	// Compare returns 1 if this > parameter
	// * returns -1 if this < parameter
	// * returns 0 if this = parameter
	Compare(Data) int

	String() string

	GetType() ID
	ZeroValue() Data
	// GenerateHash returns the hash of the Data as an ID
	// * Returns ID of empty string when the value of Data is that Data type's zero value
	GenerateHash() ID
}
