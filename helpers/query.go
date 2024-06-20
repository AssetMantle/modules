// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"net/http"

	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

type Query interface {
	GetName() string
	Command() *cobra.Command
	HandleQuery(context.Context, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(client.Context) http.HandlerFunc
	RegisterService(module.Configurator)
	RegisterGRPCGatewayRoute(client.Context, *runtime.ServeMux)
	Initialize(Mapper, ParameterManager, ...interface{}) Query
}
