package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*schema.Chain)(nil), nil)
	codec.RegisterInterface((*schema.Chains)(nil), nil)
	codec.RegisterInterface((*schema.Classification)(nil), nil)
	codec.RegisterInterface((*schema.Fact)(nil), nil)
	codec.RegisterInterface((*utility.GenesisState)(nil), nil)
	codec.RegisterInterface((*schema.Height)(nil), nil)
	codec.RegisterInterface((*schema.ID)(nil), nil)
	codec.RegisterInterface((*schema.Immutables)(nil), nil)
	codec.RegisterInterface((*schema.InterIdentity)(nil), nil)
	codec.RegisterInterface((*schema.InterIdentities)(nil), nil)
	codec.RegisterInterface((*schema.InterNFT)(nil), nil)
	codec.RegisterInterface((*schema.InterNFTs)(nil), nil)
	codec.RegisterInterface((*schema.InterNFTWallet)(nil), nil)
	codec.RegisterInterface((*schema.Maintainer)(nil), nil)
	codec.RegisterInterface((*schema.Maintainers)(nil), nil)
	codec.RegisterInterface((*schema.Mutables)(nil), nil)
	codec.RegisterInterface((*schema.NFT)(nil), nil)
	codec.RegisterInterface((*schema.NFTWallet)(nil), nil)
	codec.RegisterInterface((*schema.Properties)(nil), nil)
	codec.RegisterInterface((*schema.Property)(nil), nil)
	codec.RegisterInterface((*utility.QueryRequest)(nil), nil)
	codec.RegisterInterface((*utility.QueryResponse)(nil), nil)
	codec.RegisterInterface((*utility.Request)(nil), nil)
	codec.RegisterInterface((*schema.Share)(nil), nil)
	codec.RegisterInterface((*schema.Signature)(nil), nil)
	codec.RegisterInterface((*schema.Signatures)(nil), nil)
	codec.RegisterInterface((*schema.Split)(nil), nil)
	codec.RegisterInterface((*schema.Splits)(nil), nil)
	codec.RegisterInterface((*schema.Trait)(nil), nil)
	codec.RegisterInterface((*schema.Traits)(nil), nil)
	codec.RegisterInterface((*utility.TransactionRequest)(nil), nil)
	schema.RegisterCodec(codec)
}
