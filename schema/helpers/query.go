// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc"
)

type Query interface {
	GetName() string
	Command() *cobra.Command
	HandleMessage(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(client.Context) http.HandlerFunc
	GRPCGatewayHandler(client.Context) (method string, pattern runtime.Pattern, handlerFunc runtime.HandlerFunc)
	Service() (*grpc.ServiceDesc, interface{})
	Initialize(Mapper, Parameters, ...interface{}) Query
	GetGRPCConfigurator() GRPCConfigurator
}
