/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Query = base.NewQuery(
	mapper.ModuleName,
	QueryName,
	QueryRoute,
	QueryShort,
	QueryLong,
	packageCodec,
	registerCodec,
	initializeQueryKeeper,
	queryRequestPrototype,
	queryResponsePrototype,
	[]helpers.CLIFlag{constants.IdentityID},
)
