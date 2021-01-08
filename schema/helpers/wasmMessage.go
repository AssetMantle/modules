/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import "encoding/json"

type WasmMessage interface {
	GetType() string
	GetRawMessage() json.RawMessage
}
type WasmMessagePrototype func() WasmMessage
