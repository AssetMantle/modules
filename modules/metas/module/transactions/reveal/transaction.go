// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"context"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/cosmos/cosmos-sdk/client"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Configurator struct{}

var _ helpers.GRPCConfigurator = &Configurator{}

func (c Configurator) ConfigureGRPCServer(cfg sdkModuleTypes.Configurator) {
	RegisterTransactionServer(cfg.MsgServer(), transactionKeeper{})
}

func (c Configurator) ConfigureGRPCGatewayHandler(clientCtx client.Context, mux *runtime.ServeMux) {
	RegisterTransactionHandlerClient(context.Background(), mux, NewTransactionClient(clientCtx))
}

var Transaction = baseHelpers.NewTransaction(
	"reveal",
	"",
	"",
	requestPrototype,
	messagePrototype,
	keeperPrototype,
	Configurator{},
	constants.Data,
)
