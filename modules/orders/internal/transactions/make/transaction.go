/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"make",
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
