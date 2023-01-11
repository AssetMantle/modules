// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"

	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/AssetMantle/modules/schema/ids"
)

type Parameters interface {
	String() string

	Validate() error
	Equal(Parameters) bool

	Get(ids.ID) Parameter
	GetList() []Parameter

	Fetch(context.Context, ids.ID) Parameters
	Mutate(context.Context, Parameter) Parameters

	GetKeyTable() paramsTypes.KeyTable
	// / TODO
	// subspace.ParamSet

	Initialize(paramsTypes.Subspace) Parameters
}
