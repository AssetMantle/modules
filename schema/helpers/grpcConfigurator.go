package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type GRPCConfigurator interface {
	ConfigureGRPCServer(cfg sdkModuleTypes.Configurator)
	ConfigureGRPCGatewayHandler(clientCtx client.Context, mux *runtime.ServeMux)
}
