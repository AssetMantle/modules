package encoding

import (
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/cosmos/gogoproto/proto"
)

func Encode[Message proto.Message](message Message) ([]byte, error) {
	return base.CodecPrototype().MarshalJSON(message)
}
