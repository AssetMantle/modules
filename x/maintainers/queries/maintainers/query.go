// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainers

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	helperConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/name"
	"github.com/AssetMantle/modules/x/maintainers/constants"
)

type dummy struct{}

var Query = baseHelpers.NewQuery(
	name.GetPackageName(dummy{}),
	"",
	"",
	constants.ModuleName,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	func(server grpc.ServiceRegistrar, QueryKeeper helpers.QueryKeeper) {
		RegisterQueryServer(server, QueryKeeper.(queryKeeper))
	},
	func(clientContext client.Context, serveMux *runtime.ServeMux) error {
		return RegisterQueryHandlerClient(context.Background(), serveMux, NewQueryClient(clientContext))
	},

	helperConstants.MaintainerID,
	helperConstants.Limit,
)
