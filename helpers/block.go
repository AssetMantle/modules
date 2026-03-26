// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
)

type Block interface {
	Begin(context.Context) error
	End(context.Context) error
	Initialize(Mapper, ParameterManager, ...interface{}) Block
}
