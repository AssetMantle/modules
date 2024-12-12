// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	helperConstants "github.com/AssetMantle/modules/helpers/constants"
	"google.golang.org/grpc"
)

var Query = baseHelpers.NewQuery(
	Query_serviceDesc.ServiceName,
	Query_serviceDesc.Methods[0].MethodName,
	"",
	"",

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	func(server grpc.ServiceRegistrar, QueryKeeper helpers.QueryKeeper) {
		RegisterQueryServer(server, QueryKeeper.(queryKeeper))
	},

	helperConstants.DataID,
	helperConstants.Limit,
)
