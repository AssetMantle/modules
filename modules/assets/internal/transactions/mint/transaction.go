/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"mint",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	flags.ToID,
	flags.FromID,
	flags.ClassificationID,
	flags.ImmutableMetaProperties,
	flags.ImmutableProperties,
	flags.MutableMetaProperties,
	flags.MutableProperties,
)
