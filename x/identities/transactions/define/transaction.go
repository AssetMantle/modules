// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"google.golang.org/grpc"
)

var Transaction = baseHelpers.NewTransaction(
	Msg_serviceDesc.ServiceName,
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	func(server grpc.ServiceRegistrar, keeper helpers.TransactionKeeper) {
		RegisterMsgServer(server, keeper.(transactionKeeper))
	},

	constants.FromIdentityID,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
	constants.MutableMetaProperties,
	constants.MutableProperties,
)
