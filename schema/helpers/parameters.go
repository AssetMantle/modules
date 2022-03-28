/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Parameters interface {
	String() string

	Validate() error
	Equal(Parameters) bool

	Get(types.ID) types.Parameter
	GetList() []types.Parameter

	Fetch(sdkTypes.Context, types.ID) Parameters
	Mutate(sdkTypes.Context, types.Parameter) Parameters

	GetKeyTable() paramTypes.KeyTable
	paramTypes.ParamSet

	Initialize(paramTypes.Subspace) Parameters
}
