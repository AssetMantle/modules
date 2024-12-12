// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import "github.com/cosmos/gogoproto/proto"

type GRPCRequest interface {
	proto.Message
	Request
}
