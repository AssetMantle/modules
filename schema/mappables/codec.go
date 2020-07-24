package mappables

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*Chain)(nil), nil)
	codec.RegisterInterface((*Classification)(nil), nil)
	codec.RegisterInterface((*InterIdentity)(nil), nil)
	codec.RegisterInterface((*InterNFT)(nil), nil)
	codec.RegisterInterface((*InterNFTWallet)(nil), nil)
	codec.RegisterInterface((*Maintainer)(nil), nil)
	codec.RegisterInterface((*Split)(nil), nil)
}
