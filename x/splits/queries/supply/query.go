// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supply

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	helperConstants "github.com/AssetMantle/modules/helpers/constants"
	"google.golang.org/grpc"
)

var Query = baseHelpers.NewQuery(
	Query_serviceDesc.ServiceName,
	"",
	"",

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	func(server grpc.ServiceRegistrar, QueryKeeper helpers.QueryKeeper) {
		RegisterQueryServer(server, QueryKeeper.(queryKeeper))
	},

	helperConstants.AssetID,
)
