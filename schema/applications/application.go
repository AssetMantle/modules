/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
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
	GetLegacyAminoCodec() *codec.LegacyAmino
	GetCodec() codec.Marshaler
	GetInterfaceRegistry() types.InterfaceRegistry

	ExportApplicationStateAndValidators(bool, []string) (serverTypes.ExportedApp, error)

	Name() string
	AppVersion() string
	Logger() log.Logger
	MsgServiceRouter() *baseapp.MsgServiceRouter
	MountStores(keys ...sdkTypes.StoreKey)
	MountKVStores(keys map[string]*sdkTypes.KVStoreKey)
	MountTransientStores(keys map[string]*sdkTypes.TransientStoreKey)
	MountMemoryStores(keys map[string]*sdkTypes.MemoryStoreKey)
	MountStore(key sdkTypes.StoreKey, typ sdkTypes.StoreType)
	LoadLatestVersion() error
	LoadVersion(version int64) error
	LastCommitID() sdkTypes.CommitID
	LastBlockHeight() int64
	Router() sdkTypes.Router
	QueryRouter() sdkTypes.QueryRouter
	Seal()
	IsSealed() bool

	Initialize(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, clientTxConfig client.TxConfig, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, appOpts serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp)) Application
}
