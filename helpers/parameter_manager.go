// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/parameters"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type ParameterManager interface {
	Get() lists.ParameterList
	GetValidatableParameter(ids.PropertyID) ValidatableParameter
	GetParameter(ids.PropertyID) parameters.Parameter
	ValidateParameter(parameters.Parameter) error

	Fetch(context.Context) ParameterManager
	Set(context.Context, lists.ParameterList) ParameterManager

	GetKeyTable() paramsTypes.KeyTable
	Initialize(paramsTypes.Subspace) ParameterManager
}
