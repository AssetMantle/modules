package base

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"testing"
)

// key struct, implements helpers.Key
type testKey struct {
	name string
}

var _ helpers.Key = (*testKey)(nil)

func (t testKey) GenerateStoreKeyBytes() []byte {
	return []byte(t.name)
}

func (t testKey) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(testKey{}, "test/schema/helpers/base/testKey", nil)
}

func (t testKey) IsPartial() bool {
	if t.name == "" {
		return false
	}
	return true
}

func (t testKey) Matches(key helpers.Key) bool {
	if bytes.Equal([]byte(t.name), key.GenerateStoreKeyBytes()) {
		return true
	}
	return false
}

func NewKey(name string) helpers.Key {
	return testKey{name: name}
}

func keyPrototype() helpers.Key {
	return NewKey("")
}

// mappable struct, implements helpers.Mappable
type testMappable struct {
	id string
}

var _ helpers.Mappable = (*testMappable)(nil)

func (t testMappable) GetKey() helpers.Key {
	return NewKey(t.id)
}

func (t testMappable) RegisterCodec(c *codec.Codec) {
	c.RegisterConcrete(testMappable{}, "test/schema/helpers/base/testMappable", nil)
	return
}

func NewMappable(id string) helpers.Mappable {
	return testMappable{id: id}
}

func mappablePrototype() helpers.Mappable {
	return NewMappable("")
}

func TestMapper(t *testing.T) {
	mapper := NewMapper(keyPrototype, mappablePrototype)
	storeKey := sdkTypes.NewKVStoreKey("testHelper")
	mapper = mapper.Initialize(storeKey)
	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	require.Equal(t, nil, mapper.NewCollection(context).Get(keyPrototype()))

}
