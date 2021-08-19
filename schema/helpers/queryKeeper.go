/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type QueryKeeper interface {
	Enquire(sdkTypes.Context, QueryRequest) QueryResponse
	RegisterGRPCGatewayRoute(client.Context, *runtime.ServeMux)
	Keeper
}
