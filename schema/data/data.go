// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO URI and ID data type
type Data interface {
	GetID() ids.DataID

	String() string

	GetType() ids.ID
	ZeroValue() Data
	// GenerateHash returns the hash of the Data as an ID
	// * Returns ID of empty string when the value of Data is that Data type's zero value
	GenerateHash() ids.ID

	traits.Listable
}
