// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/AssetMantle/modules/schema/ids"
)

type ParameterManager interface {
	Get() []Parameter
	GetParameter(ids.PropertyID) Parameter
	GetValidator(ids.PropertyID) func(interface{}) error

	Fetch(context.Context) ParameterManager
	Set(context.Context, ...Parameter)

	GetKeyTable() paramsTypes.KeyTable
	RESTQueryHandler(client.Context) http.HandlerFunc
	Initialize(paramsTypes.Subspace) ParameterManager
}
