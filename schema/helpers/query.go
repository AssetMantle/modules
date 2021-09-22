/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Query interface {
	GetName() string
	Command() *cobra.Command
	//GetCommand() *cobra.Command
	HandleMessageByLegacyAmino(sdkTypes.Context, *codec.LegacyAmino, abciTypes.RequestQuery) ([]byte, error)
	HandleMessage(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(client.Context) http.HandlerFunc
	Initialize(Mapper, Parameters, ...interface{}) Query
	RegisterGRPCGatewayRoute(client.Context, *runtime.ServeMux)
	RegisterService(module.Configurator)
}
