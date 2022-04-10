// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"

	"github.com/AssetMantle/modules/schema/types"
)

type Parameters interface {
	String() string

	Validate() error
	Equal(Parameters) bool

	Get(types.ID) types.Parameter
	GetList() []types.Parameter

	Fetch(sdkTypes.Context, types.ID) Parameters
	Mutate(sdkTypes.Context, types.Parameter) Parameters

	GetKeyTable() subspace.KeyTable
	subspace.ParamSet

	Initialize(params.Subspace) Parameters
}
