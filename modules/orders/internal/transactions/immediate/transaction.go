// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"immediate",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	func(server grpc.Server, keeper helpers.TransactionKeeper) {
		RegisterTransactionServer(server, keeper.(transactionKeeper))
	},
	func(clientCtx client.Context, mux *runtime.ServeMux) error {
		return RegisterTransactionHandlerClient(context.Background(), mux, NewTransactionClient(clientCtx))
	},

	constants.FromID,
	constants.ClassificationID,
	constants.TakerID,
	constants.MakerOwnableID,
	constants.TakerOwnableID,
	constants.ExpiresIn,
	constants.MakerOwnableSplit,
	constants.TakerOwnableSplit,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
	constants.MutableMetaProperties,
	constants.MutableProperties,
)
