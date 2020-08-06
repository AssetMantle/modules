/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package kafka

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(KafkaCliCtx{}, "persistence-blockchain/KafkaCliCtx", nil)
	cdc.RegisterConcrete(KafkaMsg{}, "persistence-blockchain/KafkaMsg", nil)
}

// module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
