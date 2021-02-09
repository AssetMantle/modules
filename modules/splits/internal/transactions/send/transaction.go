/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"send",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,
	flags.FromID,
	flags.ToID,
	flags.OwnableID,
	flags.Value,
)
