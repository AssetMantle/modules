// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"net/http"

	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"
)

type Query interface {
	GetServicePath() string
	Command() *cobra.Command
	HandleQuery(context.Context, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(client.Context) http.HandlerFunc
	RegisterService(module.Configurator)
	Initialize(Mapper, ParameterManager, ...interface{}) Query
}
