// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"
)

type Transaction interface {
	GetServicePath() string
	Command() *cobra.Command
	HandleMessage(context.Context, Message) (*sdkTypes.Result, error)
	RESTRequestHandler(client.Context) http.HandlerFunc
	RegisterInterfaces(types.InterfaceRegistry)
	RegisterService(module.Configurator)
	Initialize(Mapper, ParameterManager, ...interface{}) Transaction
}
