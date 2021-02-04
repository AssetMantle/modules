/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type TestKeepers struct {
	ClassificationsKeeper helpers.AuxiliaryKeeper
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
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		ClassificationsKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers
}

func Test_Auxiliary_Keeper_Help(t *testing.T) {

	context, keepers := CreateTestInput(t)

	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))))

	classificationID := key.NewClassificationID(base.NewID(context.ChainID()), immutableProperties, mutableProperties)

	testClassificationID := key.NewClassificationID(base.NewID(context.ChainID()), base.NewProperties(), base.NewProperties())

	keepers.ClassificationsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewClassification(testClassificationID, base.NewProperties(), base.NewProperties()))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(base.NewID(classificationID.String()), nil)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(immutableProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Classification already present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(base.NewID(testClassificationID.String()), errors.EntityAlreadyExists)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(), base.NewProperties())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Max Property Count", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base.NewID("ID10"), base.NewFact(base.NewStringData("Data10"))), base.NewProperty(base.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))), base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base.NewID("ID10"), base.NewFact(base.NewStringData("Data10"))), base.NewProperty(base.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))), base.NewProperties(base.NewProperty(base.NewID("ID11"), base.NewFact(base.NewStringData("Data11"))), base.NewProperty(base.NewID("ID12"), base.NewFact(base.NewStringData("Data12"))), base.NewProperty(base.NewID("ID13"), base.NewFact(base.NewStringData("Data13"))), base.NewProperty(base.NewID("ID14"), base.NewFact(base.NewStringData("Data14"))), base.NewProperty(base.NewID("ID15"), base.NewFact(base.NewStringData("Data15"))), base.NewProperty(base.NewID("ID16"), base.NewFact(base.NewStringData("Data16"))), base.NewProperty(base.NewID("ID17"), base.NewFact(base.NewStringData("Data17"))), base.NewProperty(base.NewID("ID18"), base.NewFact(base.NewStringData("Data18"))), base.NewProperty(base.NewID("ID19"), base.NewFact(base.NewStringData("Data19"))), base.NewProperty(base.NewID("ID20"), base.NewFact(base.NewStringData("Data20")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable and Mutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errors.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data1"))), base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("Data2"))), base.NewProperty(base.NewID("ID3"), base.NewFact(base.NewStringData("Data3"))), base.NewProperty(base.NewID("ID4"), base.NewFact(base.NewStringData("Data4"))), base.NewProperty(base.NewID("ID5"), base.NewFact(base.NewStringData("Data5"))), base.NewProperty(base.NewID("ID6"), base.NewFact(base.NewStringData("Data6"))), base.NewProperty(base.NewID("ID7"), base.NewFact(base.NewStringData("Data7"))), base.NewProperty(base.NewID("ID8"), base.NewFact(base.NewStringData("Data8"))), base.NewProperty(base.NewID("ID9"), base.NewFact(base.NewStringData("Data9"))), base.NewProperty(base.NewID("ID10"), base.NewFact(base.NewStringData("Data10")))), base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data11"))), base.NewProperty(base.NewID("ID12"), base.NewFact(base.NewStringData("Data12"))), base.NewProperty(base.NewID("ID13"), base.NewFact(base.NewStringData("Data13"))), base.NewProperty(base.NewID("ID14"), base.NewFact(base.NewStringData("Data14"))), base.NewProperty(base.NewID("ID15"), base.NewFact(base.NewStringData("Data15"))), base.NewProperty(base.NewID("ID16"), base.NewFact(base.NewStringData("Data16"))), base.NewProperty(base.NewID("ID17"), base.NewFact(base.NewStringData("Data17"))), base.NewProperty(base.NewID("ID18"), base.NewFact(base.NewStringData("Data18"))), base.NewProperty(base.NewID("ID19"), base.NewFact(base.NewStringData("Data19"))), base.NewProperty(base.NewID("ID20"), base.NewFact(base.NewStringData("Data20")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
