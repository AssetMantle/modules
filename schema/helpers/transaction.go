// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

type Transaction interface {
	GetName() string
	Command() *cobra.Command
	HandleMessage(context.Context, Message) (*sdkTypes.Result, error)
	RESTRequestHandler(client.Context) http.HandlerFunc
	RegisterLegacyAminoCodec(amino *codec.LegacyAmino)
	RegisterInterfaces(types.InterfaceRegistry)
	RegisterService(module.Configurator)
	RegisterGRPCGatewayRoute(client.Context, *runtime.ServeMux)
	DecodeTransactionRequest(json.RawMessage) (sdkTypes.Msg, error)
	InitializeKeeper(Mapper, ParameterManager, ...interface{}) Transaction
}
