/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package applications

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	tendermintABCITypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintTypes "github.com/tendermint/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
)

type Application interface {
	tendermintABCITypes.Application

	GetDefaultNodeHome() string
	GetDefaultClientHome() string
	GetModuleBasicManager() sdkTypesModule.BasicManager
	GetCodec() *codec.Codec

	LoadHeight(int64) error
	ExportApplicationStateAndValidators(bool, []string) (json.RawMessage, []tendermintTypes.GenesisValidator, error)

	Name() string
	AppVersion() string
	Logger() log.Logger
	MountStores(keys ...sdk.StoreKey)
	MountKVStores(keys map[string]*sdk.KVStoreKey)
	MountTransientStores(keys map[string]*sdk.TransientStoreKey)
	MountStoreWithDB(key sdk.StoreKey, typ sdk.StoreType, db tendermintDB.DB)
	MountStore(key sdk.StoreKey, typ sdk.StoreType)
	LoadLatestVersion(baseKey *sdk.KVStoreKey) error
	LoadVersion(version int64, baseKey *sdk.KVStoreKey) error
	LastCommitID() sdk.CommitID
	LastBlockHeight() int64
	Router() sdk.Router
	QueryRouter() sdk.QueryRouter
	Seal()
	IsSealed() bool

	Initialize(logger log.Logger, db tendermintDB.DB, traceStore io.Writer, loadLatest bool, invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string, baseAppOptions ...func(*baseapp.BaseApp)) Application
}
