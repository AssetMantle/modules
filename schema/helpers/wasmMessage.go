// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import "encoding/json"

type WasmMessage interface {
	GetType() string
	GetRawMessage() json.RawMessage
}
type WasmMessagePrototype func() WasmMessage
