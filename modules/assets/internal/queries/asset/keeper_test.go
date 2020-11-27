package asset

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"testing"
)

func CreateTestInput2(t *testing.T) (sdkTypes.Context, helpers.Mapper) {
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

	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)

	return context, mapper
}

func Test_keeper_Asset(t *testing.T) {

	context, mapper := CreateTestInput2(t)

	testQueryKeeper := keeperPrototype()

	//testAssetID := base.NewID("ID")
	immutableTraits, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableTraits, Error := base.ReadProperties("burn:S|100")
	require.Equal(t, nil, Error)
	classificationID := base.NewID("ClassificationID")
	assetID := key.NewAssetID(classificationID, base.NewImmutables(immutableTraits))
	mapper.NewCollection(context).Add(mappable.NewAsset(assetID, base.NewImmutables(immutableTraits), base.NewMutables(mutableTraits)))

	testQueryRequest := newQueryRequest(assetID)
	testQueryKeeper.Initialize(mapper, nil, nil)
	require.Equal(t, 123, testQueryKeeper.Enquire(context, testQueryRequest))

}
