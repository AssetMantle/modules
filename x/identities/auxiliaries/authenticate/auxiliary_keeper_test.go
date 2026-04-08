// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import (
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/record"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	"github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	cosmosDB "github.com/cosmos/cosmos-db"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"cosmossdk.io/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"cosmossdk.io/store"
	storeMetrics "cosmossdk.io/store/metrics"
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/x/identities/parameters"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	ChainID = "testChain"
)

var (
	moduleStoreKey = storeTypes.NewKVStoreKey(constants.ModuleName)

	provisionedAddress   = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	unprovisionedAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	testIdentity         = base.PrototypeNameIdentity().ProvisionAddress(provisionedAddress)
	testIdentityID       = testIdentity.(documents.NameIdentity).GetNameIdentityID()

	AuxiliaryKeeper = auxiliaryKeeper{mapper.Prototype().Initialize(moduleStoreKey)}

	setContext = func() sdkTypes.Context {
		memDB := cosmosDB.NewMemDB()
		commitMultiStore := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
		commitMultiStore.MountStoreWithDB(moduleStoreKey, storeTypes.StoreTypeIAVL, nil)
		_ = commitMultiStore.LoadLatestVersion()
		return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())

	}

	Context = setContext()

	_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
		Add(record.NewRecord(testIdentity))
)

func Test_auxiliaryKeeper_Help(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		request helpers.AuxiliaryRequest
		want    helpers.AuxiliaryResponse
		wantErr helpers.Error
	}{
		{
			"valid request",
			func() {},
			auxiliaryRequest{
				provisionedAddress,
				testIdentityID,
			},
			newAuxiliaryResponse(),
			nil,
		},
		{
			"invalid request",
			func() {},
			auxiliaryRequest{
				unprovisionedAddress,
				testIdentityID,
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"identity not found",
			func() {},
			auxiliaryRequest{
				provisionedAddress,
				base.NewNameIdentity(baseIDs.NewStringID("not found"), baseData.NewListData()).GetNameIdentityID(),
			},
			nil,
			errorConstants.EntityNotFound,
		},
		{
			"many identities present",
			func() {
				for i := 0; i < 10000; i++ {
					_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
						Add(record.NewRecord(base.NewNameIdentity(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewListData(baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())), baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())), baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())), baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address()))))))
				}
			},
			auxiliaryRequest{
				provisionedAddress,
				testIdentityID,
			},
			newAuxiliaryResponse(),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := AuxiliaryKeeper.Help(sdkTypes.WrapSDKContext(Context), tt.request)

			assert.Equal(t, tt.want, got, "Help() got")

			if err != nil && tt.wantErr == nil || err == nil && tt.wantErr != nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("Help() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_keeperPrototype(t *testing.T) {
	got := keeperPrototype()
	assert.Equal(t, auxiliaryKeeper{}, got)
}

func Test_auxiliaryKeeper_Initialize(t *testing.T) {
	m := mapper.Prototype().Initialize(moduleStoreKey)
	pm := parameters.Prototype().Initialize(moduleStoreKey)

	keeper := keeperPrototype().Initialize(m, pm, []interface{}{})
	require.NotNil(t, keeper)
}
