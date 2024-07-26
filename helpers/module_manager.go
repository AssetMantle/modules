package helpers

import (
	"encoding/json"
	abciTypes "github.com/cometbft/cometbft/abci/types"
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

	InitGenesis(sdkTypes.Context, codec.JSONCodec, map[string]json.RawMessage) abciTypes.ResponseInitChain
	ExportGenesisForModules(sdkTypes.Context, codec.JSONCodec, []string) map[string]json.RawMessage

	RegisterServices(module.Configurator)
	RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux)
	RegisterRESTRoutes(client.Context, *mux.Router)

	RegisterInvariants(sdkTypes.InvariantRegistry)
	RegisterInterfaces(sdkCodecTypes.InterfaceRegistry)
	RegisterLegacyAminoCodec(*codec.LegacyAmino)

	SetOrderBeginBlockers(...string)
	SetOrderEndBlockers(...string)
	SetOrderInitGenesis(...string)

	BeginBlock(sdkTypes.Context, abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock
	EndBlock(sdkTypes.Context, abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock

	RunMigrations(sdkTypes.Context, module.Configurator, module.VersionMap) (module.VersionMap, error)
}
