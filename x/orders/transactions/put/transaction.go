// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/name"
)

type dummy struct{}

var Transaction = baseHelpers.NewTransaction(
	name.GetPackageName(dummy{}),
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	func(server grpc.Server, keeper helpers.TransactionKeeper) {
		RegisterServiceServer(server, keeper.(transactionKeeper))
	},
	func(clientCtx client.Context, mux *runtime.ServeMux) error {
		return RegisterServiceHandlerClient(context.Background(), mux, NewServiceClient(clientCtx))
	},

	constants.ClassificationID,
	constants.ExpiresIn,
	constants.FromIdentityID,
	constants.MakerOwnableID,
	constants.MakerOwnableSplit,
	constants.MutableMetaProperties,
	constants.MutableProperties,
	constants.TakerID,
	constants.TakerOwnableSplit,
	constants.TakerOwnableID,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
)
