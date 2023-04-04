// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/AssetMantle/schema/x/helpers/constants"
)

var Query = baseHelpers.NewQuery(
	"identities",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	func(server grpc.Server, QueryKeeper helpers.QueryKeeper) {
		RegisterServiceServer(server, QueryKeeper.(queryKeeper))
	},
	func(clientContext client.Context, serveMux *runtime.ServeMux) error {
		return RegisterServiceHandlerClient(context.Background(), serveMux, NewServiceClient(clientContext))
	},

	constants.IdentityID,
)
