// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/parameters"
)

type Parameters interface {
	String() string

	Validate() error
	Equal(Parameters) bool

	Get(ids.ID) parameters.Parameter
	GetList() []parameters.Parameter

	Fetch(sdkTypes.Context, ids.ID) Parameters
	Mutate(sdkTypes.Context, parameters.Parameter) Parameters

	GetKeyTable() subspace.KeyTable
	subspace.ParamSet

	Initialize(params.Subspace) Parameters
}
