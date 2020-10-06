/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package asset

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/assets/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Query = base.NewQuery(
	module.Name,
	Name,
	Route,
	Short,
	Long,
	packageCodec,
	registerCodec,
	initializeQueryKeeper,
	queryRequestPrototype,
	queryResponsePrototype,
	[]helpers.CLIFlag{flags.AssetID},
)
