// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"immediate",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,
	flags.ClassificationID,
	flags.FromID,
	flags.MakerOwnableSplit,
	flags.MakerOwnableID,
	flags.TakerOwnableSplit,
	flags.TakerOwnableID,
	flags.ExpiresIn,
	flags.ImmutableMetaProperties,
	flags.ImmutableProperties,
	flags.MutableMetaProperties,
	flags.MutableProperties,
)
