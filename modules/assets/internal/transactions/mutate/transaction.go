/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mutate

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"mutate",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	flags.AssetID,
	flags.FromID,
	flags.MutableMetaProperties,
	flags.MutableProperties,
)
