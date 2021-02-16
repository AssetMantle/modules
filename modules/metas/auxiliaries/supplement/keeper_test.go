/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type TestKeepers struct {
	MetasKeeper helpers.AuxiliaryKeeper
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
		MetasKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers

}

func Test_Auxiliary_Keeper_Help(t *testing.T) {

	context, keepers := CreateTestInput(t)

	heightData, _ := base.ReadHeightData("")
	decData, _ := base.ReadDecData("")

	property1 := base.NewMetaProperty(base.NewID("id1"), base.NewMetaFact(base.NewStringData("")))
	property2 := base.NewMetaProperty(base.NewID("id2"), base.NewMetaFact(heightData))
	dec, _ := sdkTypes.NewDecFromStr("123")
	property3 := base.NewMetaProperty(base.NewID("id3"), base.NewMetaFact(decData))
	property4 := base.NewMetaProperty(base.NewID("id4"), base.NewMetaFact(base.NewIDData(base.NewID(""))))
	property5 := base.NewMetaProperty(base.NewID("id5"), base.NewMetaFact(base.NewDecData(dec)))

	var metaPropertyList []types.MetaProperty
	metaPropertyList = append(metaPropertyList, property1, property2, property3, property4, property5)

	keepers.MetasKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewMeta(decData)).Add(mappable.NewMeta(base.NewDecData(dec)))

	t.Run("Positive Case", func(t *testing.T) {
		want := newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList...), nil)
		if got := keepers.MetasKeeper.Help(context, NewAuxiliaryRequest(property1.RemoveData(), property2.RemoveData(), property3.RemoveData(), property4.RemoveData(), property5.RemoveData())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
