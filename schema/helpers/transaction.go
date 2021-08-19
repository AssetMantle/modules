/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"
	"net/http"
)

type Transaction interface {
	GetName() string
	Command() *cobra.Command
	HandleMessage(sdkTypes.Context, sdkTypes.Msg) (*sdkTypes.Result, error)
	RESTRequestHandler(client.Context) http.HandlerFunc
	RegisterLegacyAminoCodec(*codec.LegacyAmino)
	RegisterInterface(registry codecTypes.InterfaceRegistry)
	DecodeTransactionRequest(json.RawMessage) (sdkTypes.Msg, error)
	InitializeKeeper(Mapper, Parameters, ...interface{}) Transaction
	RegisterService(module.Configurator)
}
