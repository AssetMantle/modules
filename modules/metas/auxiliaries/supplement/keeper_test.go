package supplement

import (
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
	"reflect"
	"testing"
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

func Test_Transfer_Aux_Keeper_Help(t *testing.T) {

	context, keepers := CreateTestInput(t)

	property1 := base.NewMetaProperty(base.NewID("id1"), base.NewMetaFact(base.NewStringData("")))
	property2 := base.NewMetaProperty(base.NewID("id2"), base.NewMetaFact(base.NewHeightData(base.NewHeight(23))))
	decSTr, _ := sdkTypes.NewDecFromStr("23")
	property3 := base.NewMetaProperty(base.NewID("id3"), base.NewMetaFact(base.NewDecData(decSTr)))
	property4 := base.NewMetaProperty(base.NewID("id4"), base.NewMetaFact(base.NewIDData(base.NewID("2344"))))

	var metaPropertyList []types.MetaProperty
	metaPropertyList = append(metaPropertyList, property1, property2, property3, property4)

	keepers.MetasKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewMeta(base.NewStringData("234"))).Add(mappable.NewMeta(base.NewHeightData(base.NewHeight(23)))).Add(mappable.NewMeta(base.NewDecData(decSTr))).Add(mappable.NewMeta(base.NewIDData(base.NewID("2344"))))

	t.Run("PositiveCase - ", func(t *testing.T) {
		want := newAuxiliaryResponse(base.NewMetaProperties(metaPropertyList), nil)
		if got := keepers.MetasKeeper.Help(context, NewAuxiliaryRequest(property1, property2, property3, property4)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
