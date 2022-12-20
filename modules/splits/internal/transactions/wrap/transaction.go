// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"wrap",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	func(grpc.Server, helpers.TransactionKeeper) {
		panic("implement me")
	},
	func(client.Context, *runtime.ServeMux) error {
		panic("implement me")
	},

	constants.FromID,
	constants.Coins,
)
