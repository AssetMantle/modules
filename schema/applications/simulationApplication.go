// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package applications

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	simAppParams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"io"
)

type SimulationApplication interface {
	simapp.App

	AppCodec() codec.Codec
	InterfaceRegistry() codecTypes.InterfaceRegistry
	GetKey(storeKey string) *sdkTypes.KVStoreKey
	GetTKey(storeKey string) *sdkTypes.TransientStoreKey
	GetMemKey(storeKey string) *sdkTypes.MemoryStoreKey
	GetSubspace(moduleName string) paramsTypes.Subspace
	SimulationManager() *module.SimulationManager
	RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig)
	RegisterTxService(clientCtx client.Context)
	RegisterTendermintService(clientCtx client.Context)

	NewSimulationApplication(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool, homePath string, invCheckPeriod uint, encodingConfig simAppParams.EncodingConfig, appOpts serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) SimulationApplication
}
