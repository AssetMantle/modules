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
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

type moduleManager struct {
	basicModules       []helpers.BasicModule
	orderInitGenesis   []string
	orderExportGenesis []string
	orderBeginBlockers []string
	orderEndBlockers   []string
	orderMigrations    []string
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
	return moduleManager.getModuleManager().InitGenesis(context, jsonCodec, genesisData)
}
func (moduleManager moduleManager) GetVersionMap() sdkModuleTypes.VersionMap {
	return moduleManager.getModuleManager().GetVersionMap()
}
func (moduleManager moduleManager) RegisterServices(configurator sdkModuleTypes.Configurator) {
	moduleManager.getModuleManager().RegisterServices(configurator)
}
func (moduleManager moduleManager) SetOrderBeginBlockers(moduleName ...string) {
	sdkModuleManager := moduleManager.getModuleManager()
	sdkModuleManager.SetOrderBeginBlockers(moduleName...)
	moduleManager.orderBeginBlockers = sdkModuleManager.OrderBeginBlockers
}
func (moduleManager moduleManager) SetOrderEndBlockers(moduleName ...string) {
	sdkModuleManager := moduleManager.getModuleManager()
	sdkModuleManager.SetOrderEndBlockers(moduleName...)
	moduleManager.orderEndBlockers = sdkModuleManager.OrderEndBlockers
}
func (moduleManager moduleManager) SetOrderInitGenesis(moduleName ...string) {
	sdkModuleManager := moduleManager.getModuleManager()
	sdkModuleManager.SetOrderInitGenesis(moduleName...)
	moduleManager.orderInitGenesis = sdkModuleManager.OrderInitGenesis
}
func (moduleManager moduleManager) BeginBlock(context sdkTypes.Context, requestBeginBlock abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return moduleManager.getModuleManager().BeginBlock(context, requestBeginBlock)
}
func (moduleManager moduleManager) EndBlock(context sdkTypes.Context, requestEndBlock abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return moduleManager.getModuleManager().EndBlock(context, requestEndBlock)
}
func (moduleManager moduleManager) RunMigrations(context sdkTypes.Context, configurator sdkModuleTypes.Configurator, versionMap sdkModuleTypes.VersionMap) (sdkModuleTypes.VersionMap, error) {
	return moduleManager.getModuleManager().RunMigrations(context, configurator, versionMap)
}
func (moduleManager moduleManager) RegisterInvariants(invariantRegistry sdkTypes.InvariantRegistry) {
	moduleManager.getModuleManager().RegisterInvariants(invariantRegistry)
}
func (moduleManager moduleManager) GetBasicManager() sdkModuleTypes.BasicManager {
	return sdkModuleTypes.NewBasicManager(moduleManager.getAppModulesBasic()...)
}
func (moduleManager moduleManager) ExportGenesisForModules(context sdkTypes.Context, jsonCodec sdkCodec.JSONCodec, moduleNames []string) map[string]json.RawMessage {
	return moduleManager.getModuleManager().ExportGenesisForModules(context, jsonCodec, moduleNames)
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
func (moduleManager moduleManager) getModuleManager() *sdkModuleTypes.Manager {
	appModules := make([]sdkModuleTypes.AppModule, len(moduleManager.basicModules))
	for i, basicModule := range moduleManager.basicModules {
		appModules[i] = basicModule
	}
	return sdkModuleTypes.NewManager(appModules...)
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
