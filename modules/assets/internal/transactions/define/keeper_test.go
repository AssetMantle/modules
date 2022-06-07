// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/constants/test"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
)

type TestKeepers struct {
	AssetsKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	scrubAuxiliary := scrub.AuxiliaryMock.Initialize(mapper, Parameters)
	defineAuxiliary := define.AuxiliaryMock.Initialize(mapper, Parameters)
	superAuxiliary := super.AuxiliaryMock.Initialize(mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(mapper, Parameters)
	keepers := TestKeepers{
		AssetsKeeper: keeperPrototype().Initialize(
			mapper,
			Parameters,
			[]interface{}{
				scrubAuxiliary,
				verifyAuxiliary,
				defineAuxiliary,
				superAuxiliary,
			}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers := CreateTestInput(t)
	immutableMetaProperties, err := utilities.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableMetaProperties, err := utilities.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)
	superMockErrorProperties, err := utilities.ReadMetaProperties("superError:S|mockError")
	require.Equal(t, nil, err)
	scrubMockErrorProperties, err := utilities.ReadMetaProperties("scrubError:S|mockError")
	require.Equal(t, nil, err)
	gt22Properties, err := utilities.ReadMetaProperties("0:S|0,1:S|1,2:S|2,3:S|3,4:S|4,5:S|5,6:S|6,7:S|7,8:S|8,9:S|9,10:S|10,11:S|11,12:S|12,13:S|13,14:S|14,15:S|15,16:S|16,17:S|17,18:S|18,19:S|19,20:S|20,21:S|21")
	require.Equal(t, nil, err)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := baseIDs.NewID("fromIdentityID")

	t.Run("PositiveCase", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(nil)
		if got := keepers.AssetsKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, immutableMetaProperties,
			immutableProperties, mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - IdentityNil", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.AssetsKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, immutableMetaProperties,
			immutableProperties, mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - ImmutableMeta Scrub Error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.AssetsKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, scrubMockErrorProperties,
			immutableProperties, mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - MutableMeta Scrub Error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.AssetsKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, immutableMetaProperties,
			immutableProperties, scrubMockErrorProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - ClassificationID define Error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.InvalidRequest)
		if got := keepers.AssetsKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, immutableMetaProperties,
			immutableProperties, mutableMetaProperties, gt22Properties.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - maintainer super Error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.AssetsKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, immutableMetaProperties,
			immutableProperties, mutableMetaProperties, superMockErrorProperties.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
