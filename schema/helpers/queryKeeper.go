/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type QueryKeeper interface {
	QueryInKeeper(sdkTypes.Context, QueryRequest) (json.RawMessage,error)
	LegacyEnquire(sdkTypes.Context, QueryRequest) QueryResponse
	RegisterGRPCGatewayRoute(client.Context, *runtime.ServeMux)
	RegisterService(module.Configurator)
	Keeper
}
