// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

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

	constants.ClassificationID,
	constants.ExpiresIn,
	constants.FromIdentityID,
	constants.MakerAssetID,
	constants.MakerSplit,
	constants.MutableMetaProperties,
	constants.MutableProperties,
	constants.TakerID,
	constants.TakerSplit,
	constants.TakerAssetID,
	constants.ImmutableMetaProperties,
	constants.ImmutableProperties,
)
