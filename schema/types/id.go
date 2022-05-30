// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/modules/schema/capabilities"
)

type ID interface {
	capabilities.Listable
	String() string
	Bytes() []byte
}
