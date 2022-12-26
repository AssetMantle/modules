// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Query = baseHelpers.NewQuery(
	"identities",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	func(server grpc.Server, keeper helpers.QueryKeeper) {
		RegisterQueryServer(server, keeper.(queryKeeper))
	},
	func(ctx client.Context, mux *runtime.ServeMux) error {
		err := RegisterQueryHandlerClient(context.Background(), mux, NewQueryClient(ctx))
		return err
	},

	constants.IdentityID,
)
