package base

import (
	"encoding/json"
	"github.com/AssetMantle/modules/helpers"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

type moduleManager struct {
	basicModules       []helpers.BasicModule
	orderInitGenesis   []string
	orderExportGenesis []string
	orderBeginBlockers []string
	orderEndBlockers   []string
}

var _ helpers.ModuleManager = (*moduleManager)(nil)

func (moduleManager moduleManager) AddTxCommands(command *cobra.Command) {
	for _, basicModule := range moduleManager.basicModules {
		if cmd := basicModule.GetTxCmd(); cmd != nil {
			command.AddCommand(cmd)
		}
	}
}
func (moduleManager moduleManager) AddQueryCommands(rootQueryCmd *cobra.Command) {
	for _, basicModule := range moduleManager.basicModules {
		if cmd := basicModule.GetQueryCmd(); cmd != nil {
			rootQueryCmd.AddCommand(cmd)
		}
	}
}
func (moduleManager moduleManager) InitGenesis(context sdkTypes.Context, jsonCodec sdkCodec.JSONCodec, genesisData map[string]json.RawMessage) abciTypes.ResponseInitChain {
	return moduleManager.getManager().InitGenesis(context, jsonCodec, genesisData)
}
func (moduleManager moduleManager) GetVersionMap() sdkModuleTypes.VersionMap {
	return moduleManager.getManager().GetVersionMap()
}
func (moduleManager moduleManager) RegisterServices(configurator sdkModuleTypes.Configurator) {
	moduleManager.getManager().RegisterServices(configurator)
}
func (moduleManager moduleManager) SetOrderBeginBlockers(moduleName ...string) helpers.ModuleManager {
	sdkModuleManager := moduleManager.getManager()
	sdkModuleManager.SetOrderBeginBlockers(moduleName...)
	moduleManager.orderBeginBlockers = sdkModuleManager.OrderBeginBlockers

	return moduleManager
}
func (moduleManager moduleManager) SetOrderEndBlockers(moduleName ...string) helpers.ModuleManager {
	sdkModuleManager := moduleManager.getManager()
	sdkModuleManager.SetOrderEndBlockers(moduleName...)
	moduleManager.orderEndBlockers = sdkModuleManager.OrderEndBlockers

	return moduleManager
}
func (moduleManager moduleManager) SetOrderInitGenesis(moduleName ...string) helpers.ModuleManager {
	sdkModuleManager := moduleManager.getManager()
	sdkModuleManager.SetOrderInitGenesis(moduleName...)
	moduleManager.orderInitGenesis = sdkModuleManager.OrderInitGenesis

	return moduleManager
}
func (moduleManager moduleManager) SetOrderExportGenesis(moduleName ...string) helpers.ModuleManager {
	sdkModuleManager := moduleManager.getManager()
	sdkModuleManager.SetOrderExportGenesis(moduleName...)
	moduleManager.orderExportGenesis = sdkModuleManager.OrderExportGenesis

	return moduleManager
}
func (moduleManager moduleManager) BeginBlock(context sdkTypes.Context, requestBeginBlock abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return moduleManager.getManager().BeginBlock(context, requestBeginBlock)
}
func (moduleManager moduleManager) EndBlock(context sdkTypes.Context, requestEndBlock abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return moduleManager.getManager().EndBlock(context, requestEndBlock)
}
func (moduleManager moduleManager) RunMigrations(context sdkTypes.Context, configurator sdkModuleTypes.Configurator, versionMap sdkModuleTypes.VersionMap) (sdkModuleTypes.VersionMap, error) {
	return moduleManager.getManager().RunMigrations(context, configurator, versionMap)
}
func (moduleManager moduleManager) RegisterInvariants(invariantRegistry sdkTypes.InvariantRegistry) {
	moduleManager.getManager().RegisterInvariants(invariantRegistry)
}
func (moduleManager moduleManager) GetBasicManager() sdkModuleTypes.BasicManager {
	return sdkModuleTypes.NewBasicManager(moduleManager.getAppModulesBasic()...)
}
func (moduleManager moduleManager) ExportGenesisForModules(context sdkTypes.Context, jsonCodec sdkCodec.JSONCodec, moduleNames []string) map[string]json.RawMessage {
	return moduleManager.getManager().ExportGenesisForModules(context, jsonCodec, moduleNames)
}
func (moduleManager moduleManager) RegisterRESTRoutes(context client.Context, router *mux.Router) {
	for _, basicModule := range moduleManager.basicModules {
		if module, ok := basicModule.(helpers.Module); ok {
			module.RegisterRESTRoutes(context, router)
		}
	}
}
func (moduleManager moduleManager) RegisterGRPCGatewayRoutes(context client.Context, serverMux *runtime.ServeMux) {
	for _, basicModule := range moduleManager.basicModules {
		basicModule.RegisterGRPCGatewayRoutes(context, serverMux)
	}
}
func (moduleManager moduleManager) RegisterLegacyAminoCodec(legacyAmino *sdkCodec.LegacyAmino) {
	for _, basicModule := range moduleManager.basicModules {
		basicModule.RegisterLegacyAminoCodec(legacyAmino)
	}
}
func (moduleManager moduleManager) RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	for _, basicModule := range moduleManager.basicModules {
		basicModule.RegisterInterfaces(interfaceRegistry)
	}
}
func (moduleManager moduleManager) getManager() *sdkModuleTypes.Manager {
	appModules := make([]sdkModuleTypes.AppModule, len(moduleManager.basicModules))
	for i, basicModule := range moduleManager.basicModules {
		appModules[i] = basicModule
	}

	manager := sdkModuleTypes.NewManager(appModules...)

	if len(moduleManager.orderInitGenesis) > 0 {
		manager.SetOrderInitGenesis(moduleManager.orderInitGenesis...)
	}
	if len(moduleManager.orderExportGenesis) > 0 {
		manager.SetOrderExportGenesis(moduleManager.orderExportGenesis...)
	}
	if len(moduleManager.orderBeginBlockers) > 0 {
		manager.SetOrderBeginBlockers(moduleManager.orderBeginBlockers...)
	}
	if len(moduleManager.orderEndBlockers) > 0 {
		manager.SetOrderEndBlockers(moduleManager.orderEndBlockers...)
	}

	return manager
}
func (moduleManager moduleManager) getAppModulesBasic() []sdkModuleTypes.AppModuleBasic {
	appModules := make([]sdkModuleTypes.AppModuleBasic, len(moduleManager.basicModules))
	for i, basicModule := range moduleManager.basicModules {
		appModules[i] = basicModule
	}
	return appModules
}

func NewModuleManager(basicModules ...helpers.BasicModule) helpers.ModuleManager {
	return moduleManager{
		basicModules: basicModules,
	}
}
