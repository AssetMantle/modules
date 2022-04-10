// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	"github.com/AssetMantle/modules/constants/flags"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Transaction = baseHelpers.NewTransaction(
	"take",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,
	flags.FromID,
	flags.TakerOwnableSplit,
	flags.OrderID,
)
