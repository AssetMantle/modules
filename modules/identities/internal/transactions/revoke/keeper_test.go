// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"reflect"
	"testing"

	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/constants/test"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/revoke"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

type TestKeepers struct {
	IdentitiesKeeper helpers.TransactionKeeper
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
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
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

	revokeAuxiliary := revoke.AuxiliaryMock.Initialize(Mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		IdentitiesKeeper: keeperPrototype().Initialize(Mapper, Parameters,
			[]interface{}{verifyAuxiliary, revokeAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers := CreateTestInput(t)

	_, err := utilities.ReadProperties("maintainedProperties:S|maintainedProperties")
	require.Equal(t, nil, err)
	_, err = utilities.ReadProperties("conformError:S|mockError")
	require.Equal(t, nil, err)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := baseIDs.NewID("fromIdentityID")
	toID := baseIDs.NewID("toID")
	toID2 := baseIDs.NewID("toID2")
	classificationID := baseIDs.NewID("ClassificationID")
	mockErrorClassification := baseIDs.NewID("revokeError")

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID, classificationID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - verify identity fail", func(t *testing.T) {
		want := newTransactionResponse(test.MockError)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, toID, classificationID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - conform mock error", func(t *testing.T) {
		want := newTransactionResponse(test.MockError)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID2, mockErrorClassification)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
