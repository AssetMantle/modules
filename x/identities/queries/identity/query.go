// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/module"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var Query = baseHelpers.NewQuery(
	"identity",
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
