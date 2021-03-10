/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package classification

import (
	"testing"

	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/common"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

func CreateTestInput(t *testing.T) sdkTypes.Context {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
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

	context := sdkTypes.NewContext(commitMultiStore, tendermintProto.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context
}

func Test_Classification_Response(t *testing.T) {
	context := CreateTestInput(t)
	collection := mapper.Prototype().NewCollection(context)

	testQueryResponse := newQueryResponse(collection, nil)
	testQueryResponseWithError := newQueryResponse(collection, errors.IncorrectFormat)

	require.Equal(t, true, testQueryResponse.IsSuccessful())
	require.Equal(t, false, testQueryResponseWithError.IsSuccessful())
	require.Equal(t, nil, testQueryResponse.GetError())
	require.Equal(t, errors.IncorrectFormat, testQueryResponseWithError.GetError())

	encodedResponse, _ := testQueryResponse.Encode()
	bytes, _ := common.Codec.MarshalJSON(testQueryResponse)
	require.Equal(t, bytes, encodedResponse)

	decodedResponse, _ := queryResponse{}.Decode(bytes)
	require.Equal(t, testQueryResponse, decodedResponse)

	decodedResponse2, _ := queryResponse{}.Decode([]byte{})
	require.Equal(t, nil, decodedResponse2)

	require.Equal(t, queryResponse{}, responsePrototype())
}
