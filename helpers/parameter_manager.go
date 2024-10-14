// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/parameters"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type ParameterManager interface {
	Get() []parameters.Parameter
	GetValidatableParameter(ids.PropertyID) ValidatableParameter
	GetParameter(ids.PropertyID) parameters.Parameter
	ValidateGenesisParameters(Genesis) error

	Fetch(context.Context) ParameterManager
	Set(context.Context, []parameters.Parameter) ParameterManager

	GetKeyTable() paramsTypes.KeyTable
	Initialize(paramsTypes.Subspace) ParameterManager
}
