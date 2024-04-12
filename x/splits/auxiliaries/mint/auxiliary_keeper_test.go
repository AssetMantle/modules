// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/mapper"
	"github.com/AssetMantle/modules/x/splits/record"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/errors"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/types"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"math/rand"
	"reflect"
	"testing"
)

const (
	ChainID       = "testChain"
	Denom         = "stake"
	GenesisSupply = 1000000000
)

type mockAuxiliaryRequest struct {
	mock.Mock
}

func (*mockAuxiliaryRequest) Validate() error {
	return nil
}

var _ helpers.AuxiliaryRequest = (*mockAuxiliaryRequest)(nil)
var (
	testSendAmount = sdkTypes.NewInt(100)

	testFromIdentity   = base.NewNameIdentity(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewListData())
	testFromIdentityID = testFromIdentity.(documents.NameIdentity).GetNameIdentityID()

	testCoinAsset   = base.NewCoinAsset(Denom)
	testCoinAssetID = testCoinAsset.GetCoinAssetID()

	encodingConfig = simapp.MakeTestEncodingConfig()

	moduleStoreKey  = sdkTypes.NewKVStoreKey(constants.ModuleName)
	AuxiliaryKeeper = auxiliaryKeeper{mapper.Prototype().Initialize(moduleStoreKey)}

	setContext = func() sdkTypes.Context {
		memDB := tendermintDB.NewMemDB()
		commitMultiStore := store.NewCommitMultiStore(memDB)
		commitMultiStore.MountStoreWithDB(moduleStoreKey, sdkTypes.StoreTypeIAVL, memDB)
		_ = commitMultiStore.LoadLatestVersion()
		return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())

	}

	Context = setContext()
)

func Test_auxiliaryKeeper_Help(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		request helpers.AuxiliaryRequest
		want    helpers.AuxiliaryResponse
		wantErr errors.Error
	}{
		{
			"valid request",
			func() {},
			auxiliaryRequest{
				OwnerID: testFromIdentityID,
				AssetID: testCoinAssetID,
				Value:   testSendAmount,
			},
			newAuxiliaryResponse(),
			nil,
		},
		{
			"invalid ownerID",
			func() {},
			auxiliaryRequest{
				OwnerID: &baseIDs.IdentityID{HashID: &baseIDs.HashID{IDBytes: []byte("invalid")}},
				AssetID: testCoinAssetID,
				Value:   testSendAmount,
			},
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"invalid assetID",
			func() {},
			auxiliaryRequest{
				OwnerID: testFromIdentityID,
				AssetID: &baseIDs.AssetID{HashID: &baseIDs.HashID{IDBytes: []byte("invalid")}},
				Value:   testSendAmount,
			},
			nil,
			errorConstants.InvalidRequest,
		}, {
			"invalid value",
			func() {},
			auxiliaryRequest{
				OwnerID: testFromIdentityID,
				AssetID: testCoinAssetID,
				Value:   sdkTypes.NewInt(-1),
			},
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"invalid request type",
			func() {},
			&mockAuxiliaryRequest{},
			nil,
			errorConstants.InvalidRequest,
		},
		{
			"with many splits",
			func() {
				for i := 0; i < 100000; i++ {
					_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
						Add(record.NewRecord(baseIDs.NewSplitID(base.NewCoinAsset(random.GenerateUniqueIdentifier()).GetCoinAssetID(), base.NewNameIdentity(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewListData()).GetNameIdentityID()), baseTypes.NewSplit(sdkTypes.NewInt(int64(rand.Intn(100000000000))))))
				}
			},
			auxiliaryRequest{
				OwnerID: testFromIdentityID,
				AssetID: testCoinAssetID,
				Value:   testSendAmount,
			},
			newAuxiliaryResponse(),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			var splitBefore types.Split
			if tt.wantErr == nil {
				splitBefore = mappable.GetSplit(AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).FetchRecord(key.NewKey(baseIDs.NewSplitID(testCoinAssetID, testFromIdentityID))).GetMappable())
			}
			got, err := AuxiliaryKeeper.Help(sdkTypes.WrapSDKContext(Context), tt.request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() got = %v, want %v", got, tt.want)
			}

			if err != nil && tt.wantErr == nil || err == nil && tt.wantErr != nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("\n want error: \n %v \n got error: \n %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				splitAfter := mappable.GetSplit(AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).FetchRecord(key.NewKey(baseIDs.NewSplitID(testCoinAssetID, testFromIdentityID))).GetMappable())
				if splitBefore == nil {
					if !splitAfter.GetValue().Equal(testSendAmount) {
						t.Errorf("incorrect split value after minting")
					}
				} else {
					if !splitAfter.GetValue().Sub(splitBefore.GetValue()).Equal(testSendAmount) {
						t.Errorf("incorrect split value after minting")
					}
				}
			}
		})
	}
}
