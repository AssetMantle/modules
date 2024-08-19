// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"google.golang.org/grpc"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	_Msg_serviceDesc.ServiceName,
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	func(server grpc.ServiceRegistrar, keeper helpers.TransactionKeeper) {
		RegisterMsgServer(server, keeper.(transactionKeeper))
	},

	constants.FromIdentityID,
	constants.ToIdentityID,
	constants.ClassificationID,
	constants.MaintainedProperties,
	constants.CanIssueIdentity,
	constants.CanQuashIdentity,
	constants.CanAddMaintainer,
	constants.CanRemoveMaintainer,
	constants.CanMutateMaintainer,
)
