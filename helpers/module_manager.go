package helpers

import (
	"encoding/json"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkCodecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

type ModuleManager interface {
	GetVersionMap() module.VersionMap
	GetBasicManager() module.BasicManager

	AddTxCommands(*cobra.Command)
	AddQueryCommands(*cobra.Command)

	InitGenesis(sdkTypes.Context, codec.JSONCodec, map[string]json.RawMessage) (*abci.ResponseInitChain, error)
	ExportGenesisForModules(sdkTypes.Context, codec.JSONCodec, []string) (map[string]json.RawMessage, error)

	RegisterServices(module.Configurator)
	RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux)
	RegisterRESTRoutes(client.Context, *mux.Router)

	RegisterInvariants(sdkTypes.InvariantRegistry)
	RegisterInterfaces(sdkCodecTypes.InterfaceRegistry)
	RegisterLegacyAminoCodec(*codec.LegacyAmino)

	SetOrderPreBlockers(...string) ModuleManager
	SetOrderBeginBlockers(...string) ModuleManager
	SetOrderEndBlockers(...string) ModuleManager
	SetOrderInitGenesis(...string) ModuleManager
	SetOrderExportGenesis(...string) ModuleManager

	PreBlock(sdkTypes.Context) (*sdkTypes.ResponsePreBlock, error)
	BeginBlock(sdkTypes.Context) (sdkTypes.BeginBlock, error)
	EndBlock(sdkTypes.Context) (sdkTypes.EndBlock, error)

	RunMigrations(sdkTypes.Context, module.Configurator, module.VersionMap) (module.VersionMap, error)
}
