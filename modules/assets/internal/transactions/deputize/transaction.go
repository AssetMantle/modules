// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

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
	"deputize",
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
	constants.ToID,
	constants.ClassificationID,
	constants.MaintainedProperties,
	constants.CanMintAsset,
	constants.CanBurnAsset,
	constants.CanRenumerateAsset,
	constants.CanAddMaintainer,
	constants.CanRemoveMaintainer,
	constants.CanMutateMaintainer,
)
