// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"

	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/parameters"
)

type Parameters interface {
	String() string

	Validate() error
	Equal(Parameters) bool

	Get(ids.ID) parameters.Parameter
	GetList() []parameters.Parameter

	Fetch(context.Context, ids.ID) Parameters
	Mutate(context.Context, parameters.Parameter) Parameters

	GetKeyTable() paramsTypes.KeyTable
	// / TODO
	// subspace.ParamSet

	Initialize(paramsTypes.Subspace) Parameters
}
