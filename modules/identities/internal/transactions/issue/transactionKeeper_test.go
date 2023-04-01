// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

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
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type TestKeepers struct {
	IssueKeeper helpers.TransactionKeeper
}

var (
	parameterManager           helpers.ParameterManager
	authenticateAuxiliary      helpers.Auxiliary
	conformAuxiliary           helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
)

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	conformAuxiliary = conform.Auxiliary.Initialize(Mapper, parameterManager)
	maintainersVerifyAuxiliary = verify.Auxiliary.Initialize(Mapper, parameterManager)
	keepers := TestKeepers{
		IssueKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}).(helpers.TransactionKeeper),
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
		in1         helpers.ParameterManager
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{Mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}},
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
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	toAddress := "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"
	toAccAddress, err := sdkTypes.AccAddressFromBech32(toAddress)
	require.Nil(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID11"), baseData.NewStringData("ImmutableData")))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewStringData("MutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	fromIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	identity.Mutate(utilities.AnyPropertyListToPropertyList(immutableMetaProperties.GetList()...)...)
	keepers.IssueKeeper.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(mappable.NewMappable(identity))
	type fields struct {
		mapper                     helpers.Mapper
		authenticateAuxiliary      helpers.Auxiliary
		conformAuxiliary           helpers.Auxiliary
		maintainersVerifyAuxiliary helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		msg     helpers.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		// NOTE: When test individually run 2nd test will fail
		{"+ve", fields{mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, toAccAddress, fromIdentityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(nil)},
		{"+ve Entity Already Exists", fields{mapper, authenticateAuxiliary, conformAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, toAccAddress, fromIdentityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(errorConstants.EntityAlreadyExists)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                     tt.fields.mapper,
				authenticateAuxiliary:      tt.fields.authenticateAuxiliary,
				conformAuxiliary:           tt.fields.conformAuxiliary,
				maintainersVerifyAuxiliary: tt.fields.maintainersVerifyAuxiliary,
			}
			if got := transactionKeeper.Transact(sdkTypes.WrapSDKContext(tt.args.context), tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
