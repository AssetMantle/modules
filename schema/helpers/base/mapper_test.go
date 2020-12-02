package base

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"testing"
)

// key struct, implements helpers.Key
type testKey struct {
	ID string
}

var _ helpers.Key = (*testKey)(nil)

func (t testKey) GenerateStoreKeyBytes() []byte {
	return append([]byte{0x11}, []byte(t.ID)...)
}

func (t testKey) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(testKey{}, "test/testKey", nil)
}

func (t testKey) IsPartial() bool {
	if t.ID == "" {
		return false
	}
	return true
}

func (t testKey) Matches(key helpers.Key) bool {
	if bytes.Equal([]byte(t.ID), key.GenerateStoreKeyBytes()) {
		return true
	}
	return false
}

func NewKey(id string) helpers.Key {
	return testKey{ID: id}
}

func keyPrototype() helpers.Key {
	return testKey{}
}

// mappable struct, implements helpers.Mappable
type testMappable struct {
	ID    string
	Value string
}

var _ helpers.Mappable = (*testMappable)(nil)

func (t testMappable) GetKey() helpers.Key {
	return NewKey(t.ID)
}

func (t testMappable) RegisterCodec(c *codec.Codec) {
	c.RegisterConcrete(testMappable{}, "test/testMappable", nil)
	return
}

func NewMappable(id string, value string) helpers.Mappable {
	return testMappable{ID: id, Value: value}
}

func mappablePrototype() helpers.Mappable {
	return testMappable{}
}

func TestMapper(t *testing.T) {

	storeKey := sdkTypes.NewKVStoreKey("test")

	mapper := NewMapper(keyPrototype, mappablePrototype).Initialize(storeKey).(mapper)

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	mapper.GetKVStoreKey()

	mapper.NewCollection(context)
	mapper.Create(context, NewMappable("test1", "value1"))
	mapper.Create(context, NewMappable("test2", "value2"))
	mapper.Read(context, NewKey("test1"))
	mapper.Update(context, NewMappable("test1", "value1"))
	mapper.Delete(context, NewKey("test2"))
	mapper.Iterate(context, NewKey("test1"), func(mappable helpers.Mappable) bool { return false })
	mapper.StoreDecoder(codec.New(), kv.Pair{
		Key: append([]byte{0x11}, []byte("test1")...), Value: mapper.codec.MustMarshalBinaryBare(NewMappable("test1", "value1"))}, kv.Pair{
		Key: append([]byte{0x11}, []byte("test1")...), Value: mapper.codec.MustMarshalBinaryBare(NewMappable("test1", "value1"))})

}
