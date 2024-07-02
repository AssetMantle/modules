// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import "github.com/cosmos/gogoproto/proto"

type QueryResponse interface {
	proto.Message
	Encode() ([]byte, error)
	Decode([]byte) (QueryResponse, error)
}
