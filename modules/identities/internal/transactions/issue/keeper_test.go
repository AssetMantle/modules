// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"fmt"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	IssueKeeper helpers.TransactionKeeper
}

var (
	Parameters                 helpers.Parameters
	authenticateAuxiliary      helpers.Auxiliary
	conformAuxiliary           helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
)

func CreateTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	types.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := types.NewKVStoreKey("test")
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := types.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters = parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, types.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := types.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)
	conformAuxiliary = conform.AuxiliaryMock.Initialize(Mapper, Parameters)
	maintainersVerifyAuxiliary = verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		IssueKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper
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
	_, _, Mapper := CreateTestInput(t)
	type fields struct {
		mapper                     helpers.Mapper
		authenticateAuxiliary      helpers.Auxiliary
		conformAuxiliary           helpers.Auxiliary
		maintainersVerifyAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		in1         helpers.Parameters
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{Mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                     tt.fields.mapper,
				authenticateAuxiliary:      tt.fields.authenticateAuxiliary,
				conformAuxiliary:           tt.fields.conformAuxiliary,
				maintainersVerifyAuxiliary: tt.fields.maintainersVerifyAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, mapper := CreateTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	toAddress := "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"
	toAccAddress, err := types.AccAddressFromBech32(toAddress)
	require.Nil(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIds.NewStringID("ID11"), baseData.NewStringData("ImmutableData")))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIds.NewStringID("authentication"), baseData.NewStringData("MutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIds.NewClassificationID(immutables, mutables)
	fromIdentityID := baseIds.NewIdentityID(classificationID, immutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]types.AccAddress{fromAccAddress}...)
	identity.Mutate(immutableMetaProperties.GetList()...)
	keepers.IssueKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(identity))
	type fields struct {
		mapper                     helpers.Mapper
		authenticateAuxiliary      helpers.Auxiliary
		conformAuxiliary           helpers.Auxiliary
		maintainersVerifyAuxiliary helpers.Auxiliary
	}
	type args struct {
		context types.Context
		msg     types.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		// NOTE: When test individually run 2nd test will fail
		{"+ve", fields{mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, toAccAddress, fromIdentityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, newTransactionResponse(nil)},
		{"+ve Entity Already Exists", fields{mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, toAccAddress, fromIdentityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, newTransactionResponse(errorConstants.EntityAlreadyExists)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                     tt.fields.mapper,
				authenticateAuxiliary:      tt.fields.authenticateAuxiliary,
				conformAuxiliary:           tt.fields.conformAuxiliary,
				maintainersVerifyAuxiliary: tt.fields.maintainersVerifyAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
