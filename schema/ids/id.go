// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	"github.com/AssetMantle/modules/schema/traits"
)

type ID interface {
	traits.Listable
	GetTypeID() StringID
	ValidateBasic() error
	FromString(string) (ID, error)
	AsString() string
	Bytes() []byte
	ToAnyID() AnyID
}
