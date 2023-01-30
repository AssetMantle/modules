package test

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/cosmos/cosmos-sdk/client"
)

var TestClientContext = client.Context{}.WithCodec(baseHelpers.CodecPrototype())
