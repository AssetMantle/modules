// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"reflect"
	"testing"

	schema "github.com/AssetMantle/schema/x"
	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/documents/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
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

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/internal/key"
	"github.com/AssetMantle/modules/x/assets/internal/mappable"
	"github.com/AssetMantle/modules/x/assets/internal/parameters"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
)

var (
	conformAuxiliary           helpers.Auxiliary
	mintAuxiliary              helpers.Auxiliary
	authenticateAuxiliary      helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
)

type TestKeepers struct {
	MintKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {
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
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

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

	mintAuxiliary = mint.Auxiliary.Initialize(Mapper, parameterManager)
	conformAuxiliary = conform.Auxiliary.Initialize(Mapper, parameterManager)
	maintainersVerifyAuxiliary = verify.Auxiliary.Initialize(Mapper, parameterManager)
	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)

	keepers := TestKeepers{
		MintKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, parameterManager
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

// func Test_transactionKeeper_Initialize(t *testing.T) {
//	_, _, mapper, parameterManager := createTestInput(t)
//	type fields struct {
//		mapper                     helpers.Mapper
//		parameterManager                 helpers.ParameterManager
//		conformAuxiliary           helpers.Auxiliary
//		mintAuxiliary              helpers.Auxiliary
//		authenticateAuxiliary      helpers.Auxiliary
//		maintainersVerifyAuxiliary helpers.Auxiliary
//	}
//	type args struct {
//		mapper      helpers.Mapper
//		parameterManager  helpers.ParameterManager
//		auxiliaries []interface{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   helpers.Keeper
//	}{
//		{"+ve", fields{Mapper, Parameters, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, Parameters, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			transactionKeeper := transactionKeeper{
//				mapper:                     tt.fields.mapper,
//				parameterManager:                 tt.fields.parameterManager,
//				conformAuxiliary:           tt.fields.conformAuxiliary,
//				mintAuxiliary:              tt.fields.mintAuxiliary,
//				authenticateAuxiliary:      tt.fields.authenticateAuxiliary,
//				maintainersVerifyAuxiliary: tt.fields.maintainersVerifyAuxiliary,
//			}
//			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
//				t.Errorf("Initialize() = %v, want %v", got, tt.want)
//			}
//		})
//	}
// }

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, mapper, parameterManager := createTestInput(t)
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testAsset := base.NewAsset(classificationID, immutables, mutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromID := baseIDs.NewIdentityID(classificationID, immutables)
	keepers.MintKeeper.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(mappable.NewMappable(testAsset))

	type fields struct {
		mapper                     helpers.Mapper
		parameterManager           helpers.ParameterManager
		conformAuxiliary           helpers.Auxiliary
		mintAuxiliary              helpers.Auxiliary
		authenticateAuxiliary      helpers.Auxiliary
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
		{"+ve", fields{mapper, parameterManager, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, fromID, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(nil)},
		{"+ve Entity Already Exists", fields{mapper, parameterManager, conformAuxiliary, mintAuxiliary, authenticateAuxiliary, maintainersVerifyAuxiliary}, args{context, newMessage(fromAccAddress, fromID, fromID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(errorConstants.EntityAlreadyExists)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                     tt.fields.mapper,
				parameterManager:           tt.fields.parameterManager,
				conformAuxiliary:           tt.fields.conformAuxiliary,
				mintAuxiliary:              tt.fields.mintAuxiliary,
				authenticateAuxiliary:      tt.fields.authenticateAuxiliary,
				maintainersVerifyAuxiliary: tt.fields.maintainersVerifyAuxiliary,
			}
			if got := transactionKeeper.Transact(sdkTypes.WrapSDKContext(context), tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
