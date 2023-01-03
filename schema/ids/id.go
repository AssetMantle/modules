// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	"github.com/AssetMantle/modules/schema/traits"
)

type ID interface {
	traits.Listable
	AsString() string
	Bytes() []byte
	ToAnyID() AnyID
}
