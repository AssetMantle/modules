// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"context"
	"github.com/AssetMantle/modules/schema/helpers"

	"github.com/cosmos/cosmos-sdk/client"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/AssetMantle/modules/modules/metas/module/module"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

type Configurator struct{}

var _ helpers.GRPCConfigurator = &Configurator{}

func (c Configurator) ConfigureGRPCServer(cfg sdkModuleTypes.Configurator) {
	RegisterQueryServer(cfg.QueryServer(), queryKeeper{})
}

func (c Configurator) ConfigureGRPCGatewayHandler(clientCtx client.Context, mux *runtime.ServeMux) {
	RegisterQueryHandlerClient(context.Background(), mux, NewQueryClient(clientCtx))
}

var Query = baseHelpers.NewQuery(
	"metas",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,
	Configurator{},

	constants.DataID,
)
