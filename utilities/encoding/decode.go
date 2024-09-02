package encoding

import (
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/cosmos/gogoproto/proto"
)

func Decode[Message proto.Message](data []byte, message Message) error {
	return base.CodecPrototype().UnmarshalJSON(data, message)
}
