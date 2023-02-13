package base

import (
	"github.com/cosmos/cosmos-sdk/client"
)

var TestClientContext = client.Context{}.WithCodec(CodecPrototype())
