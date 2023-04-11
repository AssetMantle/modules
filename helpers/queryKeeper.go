// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
)

type QueryKeeper interface {
	Enquire(context.Context, QueryRequest) QueryResponse
	Keeper
}
