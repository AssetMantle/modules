// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type QueryKeeper interface {
	QueryInKeeper(sdkTypes.Context, QueryRequest) (json.RawMessage, error)
	LegacyEnquire(sdkTypes.Context, QueryRequest) QueryResponse
	RegisterGRPCGatewayRoute(client.Context, *runtime.ServeMux)
	RegisterService(module.Configurator)
	Keeper
}
