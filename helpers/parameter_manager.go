// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/parameters"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
)

type ParameterManager interface {
	// Get returns the parameters of the parameter manager
	// NOTE: if the parameters are not fetched from the context, a nil parameter list will be returned
	Get() lists.ParameterList
	GetDefaultParameterList() lists.ParameterList
	// Set sets the parameters of the parameter manager
	// NOTE: if the parameter are not fetched from the context, it will be initialized with the default parameters, if the parameters are then updated, will be updated with the default parameters and the updated parameters
	Set(...parameters.Parameter) ParameterManager

	Validate() error

	Fetch(context.Context) ParameterManager
	Update(context.Context) (ParameterManager, error)

	Initialize(*storeTypes.KVStoreKey) ParameterManager
}
