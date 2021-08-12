/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	tendermintABCITypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type Application interface {
	tendermintABCITypes.Application

	GetDefaultHome() string
	GetModuleBasicManager() module.BasicManager
	GetLegacyAmino() *codec.LegacyAmino
	GetCodec() codec.Marshaler
	GetInterfaceRegistry() types.InterfaceRegistry

	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (servertypes.ExportedApp, error)

	Name() string
	AppVersion() string
	Logger() log.Logger
	MountStores(keys ...sdkTypes.StoreKey)
	MountKVStores(keys map[string]*sdkTypes.KVStoreKey)
	MountTransientStores(keys map[string]*sdkTypes.TransientStoreKey)
	MountStoreWithDB(key sdkTypes.StoreKey, typ sdkTypes.StoreType, db tendermintDB.DB)
	MountStore(key sdkTypes.StoreKey, typ sdkTypes.StoreType)
	LoadVersion(version int64, baseKey *sdkTypes.KVStoreKey) error
	LastCommitID() sdkTypes.CommitID
	LastBlockHeight() int64
	Router() sdkTypes.Router
	QueryRouter() sdkTypes.QueryRouter
	Seal()
	IsSealed() bool

	Initialize(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, clientTxConfig client.TxConfig, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) Application
}
