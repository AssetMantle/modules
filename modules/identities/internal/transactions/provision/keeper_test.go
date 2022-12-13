// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

type TestKeepers struct {
	ProvisionKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
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

	authenticateAuxiliary := authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		ProvisionKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{authenticateAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, Parameters
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		{"+ve", transactionKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	_, _, mapper, _parameters := CreateTestInput(t)
	supplementAuxiliary := supplement.AuxiliaryMock.Initialize(mapper, _parameters)
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		mapper helpers.Mapper
		in1    helpers.Parameters
		in2    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{mapper}, args{mapper, _parameters, []interface{}{supplementAuxiliary}}, transactionKeeper{mapper, supplementAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper: tt.fields.mapper,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, mapper, _ := CreateTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(baseLists.NewDataList()))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList()))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	toAccAddress, err := sdkTypes.AccAddressFromBech32("cosmos1u6xn6rv07p2yzzj2rm8st04x54xe5ur0t9nl5j")
	require.Nil(t, err)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutables, mutables)
	testIdentity.ProvisionAddress([]sdkTypes.AccAddress{toAccAddress}...)
	keepers.ProvisionKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(testIdentity))
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		context sdkTypes.Context
		msg     sdkTypes.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		{"+ve Not Authorized", fields{mapper}, args{context, newMessage(fromAccAddress, fromAccAddress, testFromID)}, newTransactionResponse(errorConstants.NotAuthorized)},
		{"+ve already Exists", fields{mapper}, args{context, newMessage(toAccAddress, fromAccAddress, testFromID)}, newTransactionResponse(nil)},
		{"+ve", fields{mapper}, args{context, newMessage(toAccAddress, toAccAddress, testFromID)}, newTransactionResponse(errorConstants.EntityAlreadyExists)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper: tt.fields.mapper,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
