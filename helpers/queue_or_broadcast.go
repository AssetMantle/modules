// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

// QueueOrBroadcastFunc is set by the queuing package at init time to break the
// import cycle between helpers/base and utilities/rest/queuing.
var QueueOrBroadcastFunc func(client.Context, CommonTransactionRequest, sdkTypes.Msg) error

func QueueOrBroadcastTransaction(ctx client.Context, req CommonTransactionRequest, msg sdkTypes.Msg) error {
	if QueueOrBroadcastFunc == nil {
		return fmt.Errorf("QueueOrBroadcastFunc not registered")
	}
	return QueueOrBroadcastFunc(ctx, req, msg)
}
