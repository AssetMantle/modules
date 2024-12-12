// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type Codec interface {
	client.TxConfig
	codec.Codec

	GetProtoCodec() *codec.ProtoCodec
	GetLegacyAmino() *codec.LegacyAmino
	InterfaceRegistry() types.InterfaceRegistry
	Initialize(ModuleManager) Codec
}
