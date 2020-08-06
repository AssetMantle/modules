/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package send

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	mapper.ModuleName,
	TransactionName,
	TransactionRoute,
	TransactionShort,
	TransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]helpers.CLIFlag{constants.FromID, constants.ToID, constants.OwnableID, constants.Split},
)
