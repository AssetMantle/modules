// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/revoke"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	revokeAuxiliary       helpers.Auxiliary
)

type TestKeepers struct {
	RevokeKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := types.NewKVStoreKey("test")
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := types.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, types.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)
	revokeAuxiliary = revoke.AuxiliaryMock.Initialize(Mapper, Parameters)

	context := types.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		RevokeKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
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
	_, _, Mapper, Parameters := CreateTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		parameters            helpers.Parameters
		authenticateAuxiliary helpers.Auxiliary
		revokeAuxiliary       helpers.Auxiliary
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
		{"+ve", fields{Mapper, Parameters, authenticateAuxiliary, revokeAuxiliary}, args{Mapper, Parameters, []interface{}{authenticateAuxiliary, revokeAuxiliary}}, transactionKeeper{Mapper, Parameters, authenticateAuxiliary, revokeAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameters:            tt.fields.parameters,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				revokeAuxiliary:       tt.fields.revokeAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameters, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper, Parameters := CreateTestInput(t)
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(baseLists.NewDataList())))
	immutablesMeta := baseQualified.NewImmutables(immutableMetaProperties)
	mutablesMeta := baseQualified.NewMutables(mutableMetaProperties)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutablesMeta, mutablesMeta)
	testIdentity.ProvisionAddress([]types.AccAddress{fromAccAddress}...)
	testOrder := baseDocuments.NewOrder(testClassificationID, immutablesMeta, mutablesMeta)
	keepers.RevokeKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(testOrder))
	type fields struct {
		mapper                helpers.Mapper
		parameters            helpers.Parameters
		authenticateAuxiliary helpers.Auxiliary
		revokeAuxiliary       helpers.Auxiliary
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
		{"+ve", fields{Mapper, Parameters, authenticateAuxiliary, revokeAuxiliary}, args{context, newMessage(fromAccAddress, testFromID, testFromID, testClassificationID)}, newTransactionResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameters:            tt.fields.parameters,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				revokeAuxiliary:       tt.fields.revokeAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
