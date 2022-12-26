// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classification

import (
	"context"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var Query = baseHelpers.NewQuery(
	"classifications",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	func(server grpc.Server, keeper helpers.QueryKeeper) {
		RegisterQueryServer(server, keeper.(*queryKeeper))
	},
	func(clientCtx client.Context, mux *runtime.ServeMux) error {
		return RegisterQueryHandlerClient(context.Background(), mux, NewQueryClient(clientCtx))
	},

	constants.ClassificationID,
)
