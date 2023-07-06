// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classification

import (
	"testing"

	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/key"
	"github.com/AssetMantle/modules/x/classifications/mappable"
	"github.com/AssetMantle/modules/x/classifications/parameters"
)

func CreateTestInput2(t *testing.T) (sdkTypes.Context, helpers.Keeper) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

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

	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	testQueryKeeper := keeperPrototype().Initialize(mapper, parameterManager, []interface{}{})

	return context, testQueryKeeper
}

func Test_Query_Keeper_Classification(t *testing.T) {
	context, keepers := CreateTestInput2(t)
	immutableProperties, err := base.PrototypePropertyList().FromMetaPropertiesString("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableProperties, Error2 := base.PrototypePropertyList().FromMetaPropertiesString("burn:S|100")
	require.Equal(t, nil, Error2)

	classificationID := baseIDs.NewClassificationID(baseQualified.NewImmutables(immutableProperties), baseQualified.NewMutables(mutableProperties))
	keepers.(queryKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(mappable.NewMappable(baseDocuments.NewClassification(baseQualified.NewImmutables(immutableProperties), baseQualified.NewMutables(mutableProperties))))

	testQueryRequest := newQueryRequest(classificationID)
	require.Equal(t, &QueryResponse{List: mappable.MappablesFromInterface(keepers.(queryKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(context)).Fetch(key.NewKey(classificationID)).Get())}, keepers.(queryKeeper).Enquire(sdkTypes.WrapSDKContext(context), testQueryRequest))

}
