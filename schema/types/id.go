// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/modules/schema/traits"
)

type ID interface {
	traits.Listable
	String() string
	Bytes() []byte
}
