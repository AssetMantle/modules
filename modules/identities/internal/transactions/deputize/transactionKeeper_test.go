// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"

	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/deputize"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type TestKeepers struct {
	DeputizeKeeper helpers.TransactionKeeper
}

var (
	Parameters            helpers.Parameters
	deputizeAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
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
	Parameters = parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

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

	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)
	deputizeAuxiliary = deputize.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		DeputizeKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{authenticateAuxiliary, deputizeAuxiliary}).(helpers.TransactionKeeper),
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
		mapper                helpers.Mapper
		parameters            helpers.Parameters
		deputizeAuxiliary     helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		parameters  helpers.Parameters
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{Mapper, Parameters, deputizeAuxiliary, authenticateAuxiliary}, args{Mapper, Parameters, []interface{}{deputizeAuxiliary, authenticateAuxiliary}}, transactionKeeper{Mapper, Parameters, deputizeAuxiliary, authenticateAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameters:            tt.fields.parameters,
				deputizeAuxiliary:     tt.fields.deputizeAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameters, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper := CreateTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	maintainedProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIds.NewStringID("deputize"), baseData.NewListData(baseLists.NewDataList())))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIds.NewClassificationID(immutables, mutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	fromIdentityID := baseIds.NewIdentityID(classificationID, immutables)
	toIdentityID := baseIds.NewIdentityID(classificationID, immutables)
	keepers.DeputizeKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(identity))
	type fields struct {
		mapper                helpers.Mapper
		parameters            helpers.Parameters
		deputizeAuxiliary     helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
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
		{"+ve", fields{Mapper, Parameters, deputizeAuxiliary, authenticateAuxiliary}, args{context, newMessage(fromAccAddress, fromIdentityID, toIdentityID, classificationID, maintainedProperties, true, true, true, true, true, true)}, newTransactionResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameters:            tt.fields.parameters,
				deputizeAuxiliary:     tt.fields.deputizeAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
