package helpers

type WasmMessage interface {
	Type() string
	Encode() []byte
	Decode() WasmMessage
}
