/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package liquidStaking

import (
	"github.com/persistenceOne/persistenceSDK/modules/liquidStaking/keeper"
	"github.com/persistenceOne/persistenceSDK/modules/liquidStaking/types"
)

const (
	ModuleName        = types.ModuleName
	DefaultParamSpace = types.DefaultParamSpace
	StoreKey          = types.StoreKey
)

var (
	NewKeeper       = keeper.NewKeeper
	NewGenesisState = types.NewGenesisState
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params
)
