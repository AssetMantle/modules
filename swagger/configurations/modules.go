package configurations

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmClient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeClient "github.com/cosmos/cosmos-sdk/x/upgrade/client"

	"github.com/persistenceOne/persistenceSDK/modules/assets"
	"github.com/persistenceOne/persistenceSDK/modules/classifications"
	"github.com/persistenceOne/persistenceSDK/modules/identities"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers"
	"github.com/persistenceOne/persistenceSDK/modules/metas"
	"github.com/persistenceOne/persistenceSDK/modules/orders"
	"github.com/persistenceOne/persistenceSDK/modules/splits"
	"github.com/persistenceOne/persistenceSDK/schema/applications/base/encoding"
)

var ModuleBasicManager = module.NewBasicManager(
	genutil.AppModuleBasic{},
	auth.AppModuleBasic{},
	bank.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	distribution.AppModuleBasic{},
	gov.NewAppModuleBasic(append(wasmClient.ProposalHandlers, paramsClient.ProposalHandler, distributionClient.ProposalHandler, upgradeClient.ProposalHandler, upgradeClient.CancelProposalHandler)...),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	wasm.AppModuleBasic{},
	slashing.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},

	assets.Prototype(),
	classifications.Prototype(),
	identities.Prototype(),
	maintainers.Prototype(),
	metas.Prototype(),
	orders.Prototype(),
	splits.Prototype(),
)

var EnabledWasmProposalTypeList = wasm.EnableAllProposals

func MakeEncodingConfig() encoding.EncodingConfig {
	encodingConfig := encoding.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.LegacyAmino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasicManager.RegisterLegacyAminoCodec(encodingConfig.LegacyAmino)
	ModuleBasicManager.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
