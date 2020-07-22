package base

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type wasmMessage struct {
	MsgType string          `json:"msgtype,required"`
	Raw     json.RawMessage `json:"raw,omitempty"`
}

var _ utility.WasmMessage = (*wasmMessage)(nil)

func (wasmMessage wasmMessage) Type() string                { return wasmMessage.MsgType }
func (wasmMessage wasmMessage) Encode() []byte              { return nil }
func (wasmMessage wasmMessage) Decode() utility.WasmMessage { return nil }
